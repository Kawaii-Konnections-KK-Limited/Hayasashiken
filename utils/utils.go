package utils

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/sagernet/sing/common/buf"
)

func Goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}

	return id
}

// Function to purge pools
func purgeAllocator() {

	// Reflect to get internal fields
	alloc := reflect.ValueOf(buf.DefaultAllocator)
	buffers := alloc.Elem().FieldByName("buffers")

	// Loop through pools
	for i := 0; i < buffers.Len(); i++ {
		pool := buffers.Index(i)

		// Purge pool
		pool.MethodByName("Purge").Call(nil)
	}

}
func PurgeAllocator() {

	// Reflect defaultAllocator
	if buf.DefaultAllocator == nil {
		log.Println("defaultAllocator not initialized, skipping purge")
		return
	}
	allocVal := reflect.ValueOf(buf.DefaultAllocator)

	// Check it is valid
	if !allocVal.IsValid() || allocVal.IsNil() {
		// DefaultAllocator not initialized, skip purging
		return
	}
	// Check it is a pointer
	if allocVal.Kind() != reflect.Ptr {
		return
	}

	// Dereference
	alloc := allocVal.Elem()

	// Check has buffers field
	buffersField := alloc.FieldByName("buffers")
	if !buffersField.IsValid() {
		return
	}

	// Now safe to purge buffers
	for i := 0; i < buffersField.Len(); i++ {
		pool := buffersField.Index(i)
		pool.MethodByName("Purge").Call(nil)
	}

}
