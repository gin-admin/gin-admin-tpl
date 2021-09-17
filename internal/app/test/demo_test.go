package test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-admin/gin-admin-tpl/internal/app/schema"
	"github.com/gin-admin/gin-admin-tpl/pkg/util/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	const router = apiPrefix + "v1/demos"
	var err error

	w := httptest.NewRecorder()

	// post /demos
	addItem := &schema.Demo{
		Code: uuid.MustUUID().String(),
		Name: uuid.MustUUID().String(),
	}
	engine.ServeHTTP(w, newPostRequest(router, addItem))
	assert.Equal(t, 200, w.Code)
	var addItemRes ResID
	err = parseReader(w.Body, &addItemRes)
	assert.Nil(t, err)

	// get /demos/:id
	engine.ServeHTTP(w, newGetRequest("%s/%d", nil, router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	var getItem schema.Demo
	err = parseReader(w.Body, &getItem)
	assert.Nil(t, err)
	assert.Equal(t, addItem.Name, getItem.Name)
	assert.Equal(t, addItem.Status, getItem.Status)
	assert.NotEmpty(t, getItem.ID)

	// put /demos/:id
	putItem := getItem
	putItem.Name = uuid.MustUUID().String()
	engine.ServeHTTP(w, newPutRequest("%s/%d", putItem, router, getItem.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)

	// query /demos
	engine.ServeHTTP(w, newGetRequest(router, newPageParam()))
	assert.Equal(t, 200, w.Code)
	var pageItems []*schema.Demo
	err = parsePageReader(w.Body, &pageItems)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(pageItems), 1)
	if len(pageItems) > 0 {
		assert.Equal(t, putItem.ID, pageItems[0].ID)
		assert.Equal(t, putItem.Name, pageItems[0].Name)
	}

	// delete /demos/:id
	engine.ServeHTTP(w, newDeleteRequest("%s/%d", router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)
}
