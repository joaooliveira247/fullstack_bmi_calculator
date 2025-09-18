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
		ExpectedStatus  string
		ExpectedMessage string
	}{
		{
			Body:            `{"weight": 40, "height": 1.7}`,
			ExpectedBMI:     13.84,
			ExpectedStatus:  "Severe thinness",
			ExpectedMessage: "Very low BMI ⚠️🪦 — severe health risk, seek urgent medical help.",
		},
		{
			Body:            `{"weight": 47, "height": 1.7}`,
			ExpectedBMI:     16.26,
			ExpectedStatus:  "Moderate thinness",
			ExpectedMessage: "Low BMI ⚠️🏃‍♂️ — moderate malnutrition risk, pay attention to nutrition.",
		},
		{
			Body:            `{"weight": 52, "height": 1.7}`,
			ExpectedBMI:     17.99,
			ExpectedStatus:  "Mild thinness",
			ExpectedMessage: "Slightly below normal ⚠️🍵 — improve diet to avoid deficiencies.",
		},
		{
			Body:            `{"weight": 68, "height": 1.7}`,
			ExpectedBMI:     23.53,
			ExpectedStatus:  "Normal",
			ExpectedMessage: "Healthy BMI ✅🟢 — keep up the good habits.",
		},
		{
			Body:            `{"weight": 80, "height": 1.7}`,
			ExpectedBMI:     27.68,
			ExpectedStatus:  "Overweight",
			ExpectedMessage: "Above the ideal range ⚠️🍔 — higher risk of health issues if maintained.",
		},
		{
			Body:            `{"weight": 95, "height": 1.7}`,
			ExpectedBMI:     32.87,
			ExpectedStatus:  "Obese class I",
			ExpectedMessage: "Obesity class I 🔴⚠️ — health risks present, consider professional support.",
		},
		{
			Body:            `{"weight": 110, "height": 1.7}`,
			ExpectedBMI:     38.06,
			ExpectedStatus:  "Obese class II",
			ExpectedMessage: "Obesity class II 🔴🚨 — high risk of complications, seek medical guidance.",
		},
		{
			Body:            `{"weight": 120, "height": 1.7}`,
			ExpectedBMI:     41.52,
			ExpectedStatus:  "Obese class III",
			ExpectedMessage: "Obesity class III 🛑🏥 — very high risk, medical treatment is essential.",
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
				`{"BMI": %.2f, "status": "%s", "message": "%s"}`,
				testCase.ExpectedBMI,
				testCase.ExpectedStatus,
				testCase.ExpectedMessage,
			),
			w.Body.String(),
		)
	}
}

func TestGetBMIInvalid(t *testing.T) {
	testCases := []struct {
		Body    string
		Message string
	}{
		{
			Body:    `{weight: 40, height: 1.7}`,
			Message: "invalid character 'w' looking for beginning of object key string",
		},
		{
			Body:    `{"weight": 40}`,
			Message: "Key: 'BMIIn.Height' Error:Field validation for 'Height' failed on the 'required' tag",
		},
		{
			Body:    `{"weight": "abc", "height": 1.7}`,
			Message: "json: cannot unmarshal string into Go struct field BMIIn.weight of type float64",
		},
		{
			Body:    `{}`,
			Message: `Key: 'BMIIn.Weight' Error:Field validation for 'Weight' failed on the 'required' tag\nKey: 'BMIIn.Height' Error:Field validation for 'Height' failed on the 'required' tag`,
		}}

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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(
			t,
			fmt.Sprintf(`{"message": "%s"}`, testCase.Message),
			w.Body.String(),
		)
	}
}
