package handler

import (
	"bytes"
	"errors"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/service"
	mock_service "github.com/AlibekDalgat/dynamic_segmentation/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_createSegment(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSegment, segmentInfo dynamic_segmentation.SegmentInfo)

	testTable := []struct {
		name                string
		inputBody           string
		inputSegmentInfo    dynamic_segmentation.SegmentInfo
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBidy string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"testName"}`,
			inputSegmentInfo: dynamic_segmentation.SegmentInfo{
				Name: "testName",
			},
			mockBehavior: func(s *mock_service.MockSegment, segmentInfo dynamic_segmentation.SegmentInfo) {
				s.EXPECT().CreateSegment(segmentInfo).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBidy: `{"id":1}`,
		},
		{
			name:                "equal fields",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_service.MockSegment, segmentInfo dynamic_segmentation.SegmentInfo) {},
			expectedStatusCode:  400,
			expectedRequestBidy: `{"message":"invalid input body"}`,
		},
		{
			name:      "service failed",
			inputBody: `{"name":"name"}`,
			inputSegmentInfo: dynamic_segmentation.SegmentInfo{
				Name: "name",
			},
			mockBehavior: func(s *mock_service.MockSegment, segmentInfo dynamic_segmentation.SegmentInfo) {
				s.EXPECT().CreateSegment(segmentInfo).Return(1, errors.New("service failed"))
			},
			expectedStatusCode:  500,
			expectedRequestBidy: `{"message":"service failed"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			segment := mock_service.NewMockSegment(c)
			testCase.mockBehavior(segment, testCase.inputSegmentInfo)

			services := &service.Service{Segment: segment}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/api/segment/", handler.createSegment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/segment/", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBidy, w.Body.String())
		})
	}
}
