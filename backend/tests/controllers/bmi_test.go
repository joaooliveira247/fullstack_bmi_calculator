package controllers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/fullstack_bmi_calculator/backend/src/controllers"
	"github.com/stretchr/testify/assert"
)

func TestGetBMISuccess(t *testing.T) {
	testCases := []struct {
		Body            string
		ExpectedBMI     float64
		ExpectedMessage string
	}{
		{
			Body:            `{"weight": 40, "height": 1.7}`,
			ExpectedBMI:     13.84,
			ExpectedMessage: "Severe thinness ⚠️ 🪦",
		},
		{
			Body:            `{"weight": 47, "height": 1.7}`,
			ExpectedBMI:     16.26,
			ExpectedMessage: "Modarate thinness ⚠️ 🏃‍♂️",
		},
		{
			Body:            `{"weight": 52, "height": 1.7}`,
			ExpectedBMI:     17.99,
			ExpectedMessage: "Mild thinness ⚠️ 🍵",
		},
		{
			Body:            `{"weight": 68, "height": 1.7}`,
			ExpectedBMI:     23.53,
			ExpectedMessage: "Normal ✅ 🟢",
		},
		{
			Body:            `{"weight": 80, "height": 1.7}`,
			ExpectedBMI:     27.68,
			ExpectedMessage: "Overweight ⚠️ 🍔",
		},
		{
			Body:            `{"weight": 95, "height": 1.7}`,
			ExpectedBMI:     32.87,
			ExpectedMessage: "Obese class I 🔴 ⚠️",
		},
		{
			Body:            `{"weight": 110, "height": 1.7}`,
			ExpectedBMI:     38.06,
			ExpectedMessage: "Obese class II 🔴🚨",
		},
		{
			Body:            `{"weight": 120, "height": 1.7}`,
			ExpectedBMI:     41.52,
			ExpectedMessage: "Obese class III 🛑 🏥",
		},
	}

	for _, testCase := range testCases {
		controller := controllers.NewBMIController()

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)

		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/bmi/",
			bytes.NewBufferString(testCase.Body),
		)
		c.Request.Header.Set("Content-Type", "application/json")

		controller.GetBMI(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(
			t,
			fmt.Sprintf(
				`{"BMI": %.2f, "message": "%s"}`,
				testCase.ExpectedBMI,
				testCase.ExpectedMessage,
			),
			w.Body.String(),
		)
	}
}
