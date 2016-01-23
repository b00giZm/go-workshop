package main

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"session02/kvoop/storage"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	mockStorage.EXPECT().Get("foo").Return("bar", true)

	store := New(mockStorage)

	a := assert.New(t)

	value, found := store.Get("foo");
	if !found {
		t.Fail()
	}
	a.Equal("bar", value)
}

func TestGetMissing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	mockStorage.EXPECT().Get("foo").Return(nil, false)

	store := New(mockStorage)

	if _, found := store.Get("foo"); found {
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	mockStorage.EXPECT().Set("foo", "bar").Times(1)

	store := New(mockStorage)
	store.Set("foo", "bar")
}

func TestGetMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	mockStorage.EXPECT().Get("foo").Return("bar", true)
	mockStorage.EXPECT().Get("bar").Return("baz", true)

	store := New(mockStorage)
	res := store.GetMultiple([]string{"foo", "bar"})

	a := assert.New(t)
	a.Equal(storage.KeyValueMap{"foo": "bar", "bar": "baz"}, res)
}

func TestGetMultipleWithMissing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	mockStorage.EXPECT().Get("foo").Return("bar", true)
	mockStorage.EXPECT().Get("bar").Return(nil, false)
	mockStorage.EXPECT().Get("baz").Return("meh", true)
	mockStorage.EXPECT().Get("meh").Times(0)

	store := New(mockStorage)
	res := store.GetMultiple([]string{"foo", "bar", "baz"})

	a := assert.New(t)
	a.Equal(storage.KeyValueMap{"foo": "bar", "baz": "meh"}, res)
}

func TestGetMultipleWithEmptyStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)

	store := New(mockStorage)
	mockStorage.EXPECT().Get(gomock.Any()).AnyTimes().Return(nil, false)

	res := store.GetMultiple([]string{"foo", "bar", "baz"})

	a := assert.New(t)
	a.Equal(storage.KeyValueMap{}, res)
}