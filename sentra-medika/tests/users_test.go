package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	c "sentra-medika/constants"
	m "sentra-medika/models"
	"sentra-medika/seeders"
	s "sentra-medika/services"
	u "sentra-medika/utils"
)

func SetupRouter() *gin.Engine {
	seeders.Users(u.ConnectToDatabase())
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.POST(c.Login, s.Login)

	protected := r.Group("/")
	protected.Use(s.Middleware())
	{
		admin := protected.Group("/admin")
		admin.Use(s.RoleGuard("admin"))
		{
			admin.POST(c.Users, s.CreateUser)
			admin.GET(c.Users, s.GetUsers)
			admin.PUT(c.UserByID, s.UpdateUser)
			admin.DELETE(c.UserByID, s.DeleteUser)
		}
	}

	return r
}

func GetAdminToken(r *gin.Engine) string {
	payload := map[string]string{
		"email":     "sentramedika@gmail.com",
		"password":  "admin123",
	}

	marshal, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", c.Login, bytes.NewBuffer(marshal))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		panic("Failed to obtain admin token: " + w.Body.String())
	}

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	return response["access_token"].(string)
}

func TestUserModule(t *testing.T) {
	router := SetupRouter()
	token := GetAdminToken(router)
	u.DB.Unscoped().Where("email = ?", "mizukinako7@gmail.com").Delete(&m.Users{})

	defer func() {
		u.DB.Unscoped().Where("email = ?", "mizukinako7@gmail.com").Delete(&m.Users{})
	}()

	var createdUserID string

	t.Run("Create User", func(t *testing.T) {
		payload := map[string]string{
			"full_name": "Mizuki Nakano",
			"email":     "mizukinako7@gmail.com",
			"password":  "securepassword",
			"role":      "patient",
		}

		marshal, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/admin"+c.Users, bytes.NewBuffer(marshal))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]any
		json.Unmarshal(w.Body.Bytes(), &resp)
		data := resp["data"].(map[string]any)
		createdUserID = data["ID"].(string)
	})

	t.Run("Get All Users", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/admin"+c.Users, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "mizukinako7@gmail.com")
	})

	t.Run("Update User", func(t *testing.T) {
		payload := map[string]string{
			"full_name": "Mizuki Updated",
		}

		marshal, _ := json.Marshal(payload)
		req, _ := http.NewRequest("PUT", "/admin/users/"+createdUserID, bytes.NewBuffer(marshal))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Mizuki Updated")
	})

	t.Run("Delete User", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/admin/users/"+createdUserID, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}