package to

import "testing"

func Test_Int_IntPtr(t *testing.T) {
	var i int
	var ii *int
	ii = Int_IntPtr(i)
	if IntPtr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_IntPtr")
	}
}
func Test_IntPtr_Int(t *testing.T) {
	var i *int
	var ii int
	ii = IntPtr_Int(i, UseDefaultEmpty)
	if (ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int")
	}
}

func Test_IntSlice_IntSlicePtr(t *testing.T) {
	var i []int
	var ii *[]int
	ii = IntSlice_IntSlicePtr(i)
	if len(IntSlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_IntSlicePtr")
	}
}
func Test_IntSlicePtr_IntSlice(t *testing.T) {
	var i *[]int
	var ii []int
	ii = IntSlicePtr_IntSlice(i, UseDefaultEmpty)
	if len((ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_IntSlice")
	}
}

func Test_Int_Int16(t *testing.T) {
	var i int
	var ii int16
	ii = Int_Int16(i)
	if Int16_Int(ii) != i {
		t.Errorf("error in Int_Int16")
	}
}
func Test_Int_Int16Ptr(t *testing.T) {
	var i int
	var ii *int16
	ii = Int_Int16Ptr(i)
	if Int16Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Int16Ptr")
	}
}
func Test_Int16_Int(t *testing.T) {
	var i int16
	var ii int
	ii = Int16_Int(i)
	if Int_Int16(ii) != i {
		t.Errorf("error in Int16_Int")
	}
}
func Test_Int16_IntPtr(t *testing.T) {
	var i int16
	var ii *int
	ii = Int16_IntPtr(i)
	if IntPtr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_IntPtr")
	}
}
func Test_Int16_Int16Ptr(t *testing.T) {
	var i int16
	var ii *int16
	ii = Int16_Int16Ptr(i)
	if Int16Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Int16Ptr")
	}
}
func Test_IntPtr_Int16(t *testing.T) {
	var i *int
	var ii int16
	ii = IntPtr_Int16(i, UseDefaultEmpty)
	if Int16_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int16")
	}
}
func Test_IntPtr_Int16Ptr(t *testing.T) {
	var i *int
	var ii *int16
	ii = IntPtr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int16Ptr")
	}
}
func Test_Int16Ptr_Int(t *testing.T) {
	var i *int16
	var ii int
	ii = Int16Ptr_Int(i, UseDefaultEmpty)
	if Int_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int")
	}
}
func Test_Int16Ptr_Int16(t *testing.T) {
	var i *int16
	var ii int16
	ii = Int16Ptr_Int16(i, UseDefaultEmpty)
	if (ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int16")
	}
}
func Test_Int16Ptr_IntPtr(t *testing.T) {
	var i *int16
	var ii *int
	ii = Int16Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_IntPtr")
	}
}

func Test_IntSlice_Int16Slice(t *testing.T) {
	var i []int
	var ii []int16
	ii = IntSlice_Int16Slice(i)
	if len(Int16Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Int16Slice")
	}
}
func Test_IntSlice_Int16SlicePtr(t *testing.T) {
	var i []int
	var ii *[]int16
	ii = IntSlice_Int16SlicePtr(i)
	if len(Int16SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Int16SlicePtr")
	}
}
func Test_Int16Slice_IntSlice(t *testing.T) {
	var i []int16
	var ii []int
	ii = Int16Slice_IntSlice(i)
	if len(IntSlice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_IntSlice")
	}
}
func Test_Int16Slice_IntSlicePtr(t *testing.T) {
	var i []int16
	var ii *[]int
	ii = Int16Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_IntSlicePtr")
	}
}
func Test_Int16Slice_Int16SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]int16
	ii = Int16Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Int16SlicePtr")
	}
}
func Test_IntSlicePtr_Int16Slice(t *testing.T) {
	var i *[]int
	var ii []int16
	ii = IntSlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int16Slice")
	}
}
func Test_IntSlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]int16
	ii = IntSlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_IntSlice(t *testing.T) {
	var i *[]int16
	var ii []int
	ii = Int16SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_IntSlice")
	}
}
func Test_Int16SlicePtr_Int16Slice(t *testing.T) {
	var i *[]int16
	var ii []int16
	ii = Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Int16Slice")
	}
}
func Test_Int16SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]int
	ii = Int16SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_IntSlicePtr")
	}
}

func Test_Int_Int32(t *testing.T) {
	var i int
	var ii int32
	ii = Int_Int32(i)
	if Int32_Int(ii) != i {
		t.Errorf("error in Int_Int32")
	}
}
func Test_Int_Int32Ptr(t *testing.T) {
	var i int
	var ii *int32
	ii = Int_Int32Ptr(i)
	if Int32Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Int32Ptr")
	}
}
func Test_Int32_Int(t *testing.T) {
	var i int32
	var ii int
	ii = Int32_Int(i)
	if Int_Int32(ii) != i {
		t.Errorf("error in Int32_Int")
	}
}
func Test_Int32_IntPtr(t *testing.T) {
	var i int32
	var ii *int
	ii = Int32_IntPtr(i)
	if IntPtr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_IntPtr")
	}
}
func Test_Int32_Int32Ptr(t *testing.T) {
	var i int32
	var ii *int32
	ii = Int32_Int32Ptr(i)
	if Int32Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Int32Ptr")
	}
}
func Test_IntPtr_Int32(t *testing.T) {
	var i *int
	var ii int32
	ii = IntPtr_Int32(i, UseDefaultEmpty)
	if Int32_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int32")
	}
}
func Test_IntPtr_Int32Ptr(t *testing.T) {
	var i *int
	var ii *int32
	ii = IntPtr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int32Ptr")
	}
}
func Test_Int32Ptr_Int(t *testing.T) {
	var i *int32
	var ii int
	ii = Int32Ptr_Int(i, UseDefaultEmpty)
	if Int_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int")
	}
}
func Test_Int32Ptr_Int32(t *testing.T) {
	var i *int32
	var ii int32
	ii = Int32Ptr_Int32(i, UseDefaultEmpty)
	if (ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int32")
	}
}
func Test_Int32Ptr_IntPtr(t *testing.T) {
	var i *int32
	var ii *int
	ii = Int32Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_IntPtr")
	}
}

func Test_IntSlice_Int32Slice(t *testing.T) {
	var i []int
	var ii []int32
	ii = IntSlice_Int32Slice(i)
	if len(Int32Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Int32Slice")
	}
}
func Test_IntSlice_Int32SlicePtr(t *testing.T) {
	var i []int
	var ii *[]int32
	ii = IntSlice_Int32SlicePtr(i)
	if len(Int32SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Int32SlicePtr")
	}
}
func Test_Int32Slice_IntSlice(t *testing.T) {
	var i []int32
	var ii []int
	ii = Int32Slice_IntSlice(i)
	if len(IntSlice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_IntSlice")
	}
}
func Test_Int32Slice_IntSlicePtr(t *testing.T) {
	var i []int32
	var ii *[]int
	ii = Int32Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_IntSlicePtr")
	}
}
func Test_Int32Slice_Int32SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]int32
	ii = Int32Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Int32SlicePtr")
	}
}
func Test_IntSlicePtr_Int32Slice(t *testing.T) {
	var i *[]int
	var ii []int32
	ii = IntSlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int32Slice")
	}
}
func Test_IntSlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]int32
	ii = IntSlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_IntSlice(t *testing.T) {
	var i *[]int32
	var ii []int
	ii = Int32SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_IntSlice")
	}
}
func Test_Int32SlicePtr_Int32Slice(t *testing.T) {
	var i *[]int32
	var ii []int32
	ii = Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Int32Slice")
	}
}
func Test_Int32SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]int
	ii = Int32SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_IntSlicePtr")
	}
}

func Test_Int_Int64(t *testing.T) {
	var i int
	var ii int64
	ii = Int_Int64(i)
	if Int64_Int(ii) != i {
		t.Errorf("error in Int_Int64")
	}
}
func Test_Int_Int64Ptr(t *testing.T) {
	var i int
	var ii *int64
	ii = Int_Int64Ptr(i)
	if Int64Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Int64Ptr")
	}
}
func Test_Int64_Int(t *testing.T) {
	var i int64
	var ii int
	ii = Int64_Int(i)
	if Int_Int64(ii) != i {
		t.Errorf("error in Int64_Int")
	}
}
func Test_Int64_IntPtr(t *testing.T) {
	var i int64
	var ii *int
	ii = Int64_IntPtr(i)
	if IntPtr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_IntPtr")
	}
}
func Test_Int64_Int64Ptr(t *testing.T) {
	var i int64
	var ii *int64
	ii = Int64_Int64Ptr(i)
	if Int64Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Int64Ptr")
	}
}
func Test_IntPtr_Int64(t *testing.T) {
	var i *int
	var ii int64
	ii = IntPtr_Int64(i, UseDefaultEmpty)
	if Int64_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int64")
	}
}
func Test_IntPtr_Int64Ptr(t *testing.T) {
	var i *int
	var ii *int64
	ii = IntPtr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int64Ptr")
	}
}
func Test_Int64Ptr_Int(t *testing.T) {
	var i *int64
	var ii int
	ii = Int64Ptr_Int(i, UseDefaultEmpty)
	if Int_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int")
	}
}
func Test_Int64Ptr_Int64(t *testing.T) {
	var i *int64
	var ii int64
	ii = Int64Ptr_Int64(i, UseDefaultEmpty)
	if (ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int64")
	}
}
func Test_Int64Ptr_IntPtr(t *testing.T) {
	var i *int64
	var ii *int
	ii = Int64Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_IntPtr")
	}
}

func Test_IntSlice_Int64Slice(t *testing.T) {
	var i []int
	var ii []int64
	ii = IntSlice_Int64Slice(i)
	if len(Int64Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Int64Slice")
	}
}
func Test_IntSlice_Int64SlicePtr(t *testing.T) {
	var i []int
	var ii *[]int64
	ii = IntSlice_Int64SlicePtr(i)
	if len(Int64SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Int64SlicePtr")
	}
}
func Test_Int64Slice_IntSlice(t *testing.T) {
	var i []int64
	var ii []int
	ii = Int64Slice_IntSlice(i)
	if len(IntSlice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_IntSlice")
	}
}
func Test_Int64Slice_IntSlicePtr(t *testing.T) {
	var i []int64
	var ii *[]int
	ii = Int64Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_IntSlicePtr")
	}
}
func Test_Int64Slice_Int64SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]int64
	ii = Int64Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Int64SlicePtr")
	}
}
func Test_IntSlicePtr_Int64Slice(t *testing.T) {
	var i *[]int
	var ii []int64
	ii = IntSlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int64Slice")
	}
}
func Test_IntSlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]int64
	ii = IntSlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_IntSlice(t *testing.T) {
	var i *[]int64
	var ii []int
	ii = Int64SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_IntSlice")
	}
}
func Test_Int64SlicePtr_Int64Slice(t *testing.T) {
	var i *[]int64
	var ii []int64
	ii = Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Int64Slice")
	}
}
func Test_Int64SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]int
	ii = Int64SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_IntSlicePtr")
	}
}

func Test_Int_Int8(t *testing.T) {
	var i int
	var ii int8
	ii = Int_Int8(i)
	if Int8_Int(ii) != i {
		t.Errorf("error in Int_Int8")
	}
}
func Test_Int_Int8Ptr(t *testing.T) {
	var i int
	var ii *int8
	ii = Int_Int8Ptr(i)
	if Int8Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Int8Ptr")
	}
}
func Test_Int8_Int(t *testing.T) {
	var i int8
	var ii int
	ii = Int8_Int(i)
	if Int_Int8(ii) != i {
		t.Errorf("error in Int8_Int")
	}
}
func Test_Int8_IntPtr(t *testing.T) {
	var i int8
	var ii *int
	ii = Int8_IntPtr(i)
	if IntPtr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_IntPtr")
	}
}
func Test_Int8_Int8Ptr(t *testing.T) {
	var i int8
	var ii *int8
	ii = Int8_Int8Ptr(i)
	if Int8Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Int8Ptr")
	}
}
func Test_IntPtr_Int8(t *testing.T) {
	var i *int
	var ii int8
	ii = IntPtr_Int8(i, UseDefaultEmpty)
	if Int8_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int8")
	}
}
func Test_IntPtr_Int8Ptr(t *testing.T) {
	var i *int
	var ii *int8
	ii = IntPtr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Int8Ptr")
	}
}
func Test_Int8Ptr_Int(t *testing.T) {
	var i *int8
	var ii int
	ii = Int8Ptr_Int(i, UseDefaultEmpty)
	if Int_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int")
	}
}
func Test_Int8Ptr_Int8(t *testing.T) {
	var i *int8
	var ii int8
	ii = Int8Ptr_Int8(i, UseDefaultEmpty)
	if (ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int8")
	}
}
func Test_Int8Ptr_IntPtr(t *testing.T) {
	var i *int8
	var ii *int
	ii = Int8Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_IntPtr")
	}
}

func Test_IntSlice_Int8Slice(t *testing.T) {
	var i []int
	var ii []int8
	ii = IntSlice_Int8Slice(i)
	if len(Int8Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Int8Slice")
	}
}
func Test_IntSlice_Int8SlicePtr(t *testing.T) {
	var i []int
	var ii *[]int8
	ii = IntSlice_Int8SlicePtr(i)
	if len(Int8SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Int8SlicePtr")
	}
}
func Test_Int8Slice_IntSlice(t *testing.T) {
	var i []int8
	var ii []int
	ii = Int8Slice_IntSlice(i)
	if len(IntSlice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_IntSlice")
	}
}
func Test_Int8Slice_IntSlicePtr(t *testing.T) {
	var i []int8
	var ii *[]int
	ii = Int8Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_IntSlicePtr")
	}
}
func Test_Int8Slice_Int8SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]int8
	ii = Int8Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Int8SlicePtr")
	}
}
func Test_IntSlicePtr_Int8Slice(t *testing.T) {
	var i *[]int
	var ii []int8
	ii = IntSlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int8Slice")
	}
}
func Test_IntSlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]int8
	ii = IntSlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_IntSlice(t *testing.T) {
	var i *[]int8
	var ii []int
	ii = Int8SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_IntSlice")
	}
}
func Test_Int8SlicePtr_Int8Slice(t *testing.T) {
	var i *[]int8
	var ii []int8
	ii = Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Int8Slice")
	}
}
func Test_Int8SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]int
	ii = Int8SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_IntSlicePtr")
	}
}

func Test_Int_Uint(t *testing.T) {
	var i int
	var ii uint
	ii = Int_Uint(i)
	if Uint_Int(ii) != i {
		t.Errorf("error in Int_Uint")
	}
}
func Test_Int_UintPtr(t *testing.T) {
	var i int
	var ii *uint
	ii = Int_UintPtr(i)
	if UintPtr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_UintPtr")
	}
}
func Test_Uint_Int(t *testing.T) {
	var i uint
	var ii int
	ii = Uint_Int(i)
	if Int_Uint(ii) != i {
		t.Errorf("error in Uint_Int")
	}
}
func Test_Uint_IntPtr(t *testing.T) {
	var i uint
	var ii *int
	ii = Uint_IntPtr(i)
	if IntPtr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_IntPtr")
	}
}
func Test_Uint_UintPtr(t *testing.T) {
	var i uint
	var ii *uint
	ii = Uint_UintPtr(i)
	if UintPtr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_UintPtr")
	}
}
func Test_IntPtr_Uint(t *testing.T) {
	var i *int
	var ii uint
	ii = IntPtr_Uint(i, UseDefaultEmpty)
	if Uint_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint")
	}
}
func Test_IntPtr_UintPtr(t *testing.T) {
	var i *int
	var ii *uint
	ii = IntPtr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_UintPtr")
	}
}
func Test_UintPtr_Int(t *testing.T) {
	var i *uint
	var ii int
	ii = UintPtr_Int(i, UseDefaultEmpty)
	if Int_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int")
	}
}
func Test_UintPtr_Uint(t *testing.T) {
	var i *uint
	var ii uint
	ii = UintPtr_Uint(i, UseDefaultEmpty)
	if (ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint")
	}
}
func Test_UintPtr_IntPtr(t *testing.T) {
	var i *uint
	var ii *int
	ii = UintPtr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_IntPtr")
	}
}

func Test_IntSlice_UintSlice(t *testing.T) {
	var i []int
	var ii []uint
	ii = IntSlice_UintSlice(i)
	if len(UintSlice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_UintSlice")
	}
}
func Test_IntSlice_UintSlicePtr(t *testing.T) {
	var i []int
	var ii *[]uint
	ii = IntSlice_UintSlicePtr(i)
	if len(UintSlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_UintSlicePtr")
	}
}
func Test_UintSlice_IntSlice(t *testing.T) {
	var i []uint
	var ii []int
	ii = UintSlice_IntSlice(i)
	if len(IntSlice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_IntSlice")
	}
}
func Test_UintSlice_IntSlicePtr(t *testing.T) {
	var i []uint
	var ii *[]int
	ii = UintSlice_IntSlicePtr(i)
	if len(IntSlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_IntSlicePtr")
	}
}
func Test_UintSlice_UintSlicePtr(t *testing.T) {
	var i []uint
	var ii *[]uint
	ii = UintSlice_UintSlicePtr(i)
	if len(UintSlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_UintSlicePtr")
	}
}
func Test_IntSlicePtr_UintSlice(t *testing.T) {
	var i *[]int
	var ii []uint
	ii = IntSlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_UintSlice")
	}
}
func Test_IntSlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]uint
	ii = IntSlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_UintSlicePtr")
	}
}
func Test_UintSlicePtr_IntSlice(t *testing.T) {
	var i *[]uint
	var ii []int
	ii = UintSlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_IntSlice")
	}
}
func Test_UintSlicePtr_UintSlice(t *testing.T) {
	var i *[]uint
	var ii []uint
	ii = UintSlicePtr_UintSlice(i, UseDefaultEmpty)
	if len((ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_UintSlice")
	}
}
func Test_UintSlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]int
	ii = UintSlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_IntSlicePtr")
	}
}

func Test_Int_Uint16(t *testing.T) {
	var i int
	var ii uint16
	ii = Int_Uint16(i)
	if Uint16_Int(ii) != i {
		t.Errorf("error in Int_Uint16")
	}
}
func Test_Int_Uint16Ptr(t *testing.T) {
	var i int
	var ii *uint16
	ii = Int_Uint16Ptr(i)
	if Uint16Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Uint16Ptr")
	}
}
func Test_Uint16_Int(t *testing.T) {
	var i uint16
	var ii int
	ii = Uint16_Int(i)
	if Int_Uint16(ii) != i {
		t.Errorf("error in Uint16_Int")
	}
}
func Test_Uint16_IntPtr(t *testing.T) {
	var i uint16
	var ii *int
	ii = Uint16_IntPtr(i)
	if IntPtr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_IntPtr")
	}
}
func Test_Uint16_Uint16Ptr(t *testing.T) {
	var i uint16
	var ii *uint16
	ii = Uint16_Uint16Ptr(i)
	if Uint16Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Uint16Ptr")
	}
}
func Test_IntPtr_Uint16(t *testing.T) {
	var i *int
	var ii uint16
	ii = IntPtr_Uint16(i, UseDefaultEmpty)
	if Uint16_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint16")
	}
}
func Test_IntPtr_Uint16Ptr(t *testing.T) {
	var i *int
	var ii *uint16
	ii = IntPtr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Int(t *testing.T) {
	var i *uint16
	var ii int
	ii = Uint16Ptr_Int(i, UseDefaultEmpty)
	if Int_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int")
	}
}
func Test_Uint16Ptr_Uint16(t *testing.T) {
	var i *uint16
	var ii uint16
	ii = Uint16Ptr_Uint16(i, UseDefaultEmpty)
	if (ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint16")
	}
}
func Test_Uint16Ptr_IntPtr(t *testing.T) {
	var i *uint16
	var ii *int
	ii = Uint16Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_IntPtr")
	}
}

func Test_IntSlice_Uint16Slice(t *testing.T) {
	var i []int
	var ii []uint16
	ii = IntSlice_Uint16Slice(i)
	if len(Uint16Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Uint16Slice")
	}
}
func Test_IntSlice_Uint16SlicePtr(t *testing.T) {
	var i []int
	var ii *[]uint16
	ii = IntSlice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Uint16SlicePtr")
	}
}
func Test_Uint16Slice_IntSlice(t *testing.T) {
	var i []uint16
	var ii []int
	ii = Uint16Slice_IntSlice(i)
	if len(IntSlice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_IntSlice")
	}
}
func Test_Uint16Slice_IntSlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]int
	ii = Uint16Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_IntSlicePtr")
	}
}
func Test_Uint16Slice_Uint16SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]uint16
	ii = Uint16Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Uint16SlicePtr")
	}
}
func Test_IntSlicePtr_Uint16Slice(t *testing.T) {
	var i *[]int
	var ii []uint16
	ii = IntSlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint16Slice")
	}
}
func Test_IntSlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]uint16
	ii = IntSlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_IntSlice(t *testing.T) {
	var i *[]uint16
	var ii []int
	ii = Uint16SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_IntSlice")
	}
}
func Test_Uint16SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]uint16
	var ii []uint16
	ii = Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Uint16Slice")
	}
}
func Test_Uint16SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]int
	ii = Uint16SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_IntSlicePtr")
	}
}

func Test_Int_Uint32(t *testing.T) {
	var i int
	var ii uint32
	ii = Int_Uint32(i)
	if Uint32_Int(ii) != i {
		t.Errorf("error in Int_Uint32")
	}
}
func Test_Int_Uint32Ptr(t *testing.T) {
	var i int
	var ii *uint32
	ii = Int_Uint32Ptr(i)
	if Uint32Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Uint32Ptr")
	}
}
func Test_Uint32_Int(t *testing.T) {
	var i uint32
	var ii int
	ii = Uint32_Int(i)
	if Int_Uint32(ii) != i {
		t.Errorf("error in Uint32_Int")
	}
}
func Test_Uint32_IntPtr(t *testing.T) {
	var i uint32
	var ii *int
	ii = Uint32_IntPtr(i)
	if IntPtr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_IntPtr")
	}
}
func Test_Uint32_Uint32Ptr(t *testing.T) {
	var i uint32
	var ii *uint32
	ii = Uint32_Uint32Ptr(i)
	if Uint32Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Uint32Ptr")
	}
}
func Test_IntPtr_Uint32(t *testing.T) {
	var i *int
	var ii uint32
	ii = IntPtr_Uint32(i, UseDefaultEmpty)
	if Uint32_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint32")
	}
}
func Test_IntPtr_Uint32Ptr(t *testing.T) {
	var i *int
	var ii *uint32
	ii = IntPtr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Int(t *testing.T) {
	var i *uint32
	var ii int
	ii = Uint32Ptr_Int(i, UseDefaultEmpty)
	if Int_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int")
	}
}
func Test_Uint32Ptr_Uint32(t *testing.T) {
	var i *uint32
	var ii uint32
	ii = Uint32Ptr_Uint32(i, UseDefaultEmpty)
	if (ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint32")
	}
}
func Test_Uint32Ptr_IntPtr(t *testing.T) {
	var i *uint32
	var ii *int
	ii = Uint32Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_IntPtr")
	}
}

func Test_IntSlice_Uint32Slice(t *testing.T) {
	var i []int
	var ii []uint32
	ii = IntSlice_Uint32Slice(i)
	if len(Uint32Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Uint32Slice")
	}
}
func Test_IntSlice_Uint32SlicePtr(t *testing.T) {
	var i []int
	var ii *[]uint32
	ii = IntSlice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Uint32SlicePtr")
	}
}
func Test_Uint32Slice_IntSlice(t *testing.T) {
	var i []uint32
	var ii []int
	ii = Uint32Slice_IntSlice(i)
	if len(IntSlice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_IntSlice")
	}
}
func Test_Uint32Slice_IntSlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]int
	ii = Uint32Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_IntSlicePtr")
	}
}
func Test_Uint32Slice_Uint32SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]uint32
	ii = Uint32Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Uint32SlicePtr")
	}
}
func Test_IntSlicePtr_Uint32Slice(t *testing.T) {
	var i *[]int
	var ii []uint32
	ii = IntSlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint32Slice")
	}
}
func Test_IntSlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]uint32
	ii = IntSlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_IntSlice(t *testing.T) {
	var i *[]uint32
	var ii []int
	ii = Uint32SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_IntSlice")
	}
}
func Test_Uint32SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]uint32
	var ii []uint32
	ii = Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Uint32Slice")
	}
}
func Test_Uint32SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]int
	ii = Uint32SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_IntSlicePtr")
	}
}

func Test_Int_Uint64(t *testing.T) {
	var i int
	var ii uint64
	ii = Int_Uint64(i)
	if Uint64_Int(ii) != i {
		t.Errorf("error in Int_Uint64")
	}
}
func Test_Int_Uint64Ptr(t *testing.T) {
	var i int
	var ii *uint64
	ii = Int_Uint64Ptr(i)
	if Uint64Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Uint64Ptr")
	}
}
func Test_Uint64_Int(t *testing.T) {
	var i uint64
	var ii int
	ii = Uint64_Int(i)
	if Int_Uint64(ii) != i {
		t.Errorf("error in Uint64_Int")
	}
}
func Test_Uint64_IntPtr(t *testing.T) {
	var i uint64
	var ii *int
	ii = Uint64_IntPtr(i)
	if IntPtr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_IntPtr")
	}
}
func Test_Uint64_Uint64Ptr(t *testing.T) {
	var i uint64
	var ii *uint64
	ii = Uint64_Uint64Ptr(i)
	if Uint64Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Uint64Ptr")
	}
}
func Test_IntPtr_Uint64(t *testing.T) {
	var i *int
	var ii uint64
	ii = IntPtr_Uint64(i, UseDefaultEmpty)
	if Uint64_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint64")
	}
}
func Test_IntPtr_Uint64Ptr(t *testing.T) {
	var i *int
	var ii *uint64
	ii = IntPtr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Int(t *testing.T) {
	var i *uint64
	var ii int
	ii = Uint64Ptr_Int(i, UseDefaultEmpty)
	if Int_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int")
	}
}
func Test_Uint64Ptr_Uint64(t *testing.T) {
	var i *uint64
	var ii uint64
	ii = Uint64Ptr_Uint64(i, UseDefaultEmpty)
	if (ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint64")
	}
}
func Test_Uint64Ptr_IntPtr(t *testing.T) {
	var i *uint64
	var ii *int
	ii = Uint64Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_IntPtr")
	}
}

func Test_IntSlice_Uint64Slice(t *testing.T) {
	var i []int
	var ii []uint64
	ii = IntSlice_Uint64Slice(i)
	if len(Uint64Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Uint64Slice")
	}
}
func Test_IntSlice_Uint64SlicePtr(t *testing.T) {
	var i []int
	var ii *[]uint64
	ii = IntSlice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_IntSlice(t *testing.T) {
	var i []uint64
	var ii []int
	ii = Uint64Slice_IntSlice(i)
	if len(IntSlice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_IntSlice")
	}
}
func Test_Uint64Slice_IntSlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]int
	ii = Uint64Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_IntSlicePtr")
	}
}
func Test_Uint64Slice_Uint64SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]uint64
	ii = Uint64Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Uint64SlicePtr")
	}
}
func Test_IntSlicePtr_Uint64Slice(t *testing.T) {
	var i *[]int
	var ii []uint64
	ii = IntSlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint64Slice")
	}
}
func Test_IntSlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]uint64
	ii = IntSlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_IntSlice(t *testing.T) {
	var i *[]uint64
	var ii []int
	ii = Uint64SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_IntSlice")
	}
}
func Test_Uint64SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]uint64
	var ii []uint64
	ii = Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Uint64Slice")
	}
}
func Test_Uint64SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]int
	ii = Uint64SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_IntSlicePtr")
	}
}

func Test_Int_Uint8(t *testing.T) {
	var i int
	var ii uint8
	ii = Int_Uint8(i)
	if Uint8_Int(ii) != i {
		t.Errorf("error in Int_Uint8")
	}
}
func Test_Int_Uint8Ptr(t *testing.T) {
	var i int
	var ii *uint8
	ii = Int_Uint8Ptr(i)
	if Uint8Ptr_Int(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int_Uint8Ptr")
	}
}
func Test_Uint8_Int(t *testing.T) {
	var i uint8
	var ii int
	ii = Uint8_Int(i)
	if Int_Uint8(ii) != i {
		t.Errorf("error in Uint8_Int")
	}
}
func Test_Uint8_IntPtr(t *testing.T) {
	var i uint8
	var ii *int
	ii = Uint8_IntPtr(i)
	if IntPtr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_IntPtr")
	}
}
func Test_Uint8_Uint8Ptr(t *testing.T) {
	var i uint8
	var ii *uint8
	ii = Uint8_Uint8Ptr(i)
	if Uint8Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Uint8Ptr")
	}
}
func Test_IntPtr_Uint8(t *testing.T) {
	var i *int
	var ii uint8
	ii = IntPtr_Uint8(i, UseDefaultEmpty)
	if Uint8_Int(ii) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint8")
	}
}
func Test_IntPtr_Uint8Ptr(t *testing.T) {
	var i *int
	var ii *uint8
	ii = IntPtr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Int(ii, UseDefaultEmpty) != IntPtr_Int(i, UseDefaultEmpty) {
		t.Errorf("error in IntPtr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Int(t *testing.T) {
	var i *uint8
	var ii int
	ii = Uint8Ptr_Int(i, UseDefaultEmpty)
	if Int_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int")
	}
}
func Test_Uint8Ptr_Uint8(t *testing.T) {
	var i *uint8
	var ii uint8
	ii = Uint8Ptr_Uint8(i, UseDefaultEmpty)
	if (ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint8")
	}
}
func Test_Uint8Ptr_IntPtr(t *testing.T) {
	var i *uint8
	var ii *int
	ii = Uint8Ptr_IntPtr(i, UseDefaultEmpty)
	if IntPtr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_IntPtr")
	}
}

func Test_IntSlice_Uint8Slice(t *testing.T) {
	var i []int
	var ii []uint8
	ii = IntSlice_Uint8Slice(i)
	if len(Uint8Slice_IntSlice(ii)) != len(i) {
		t.Errorf("error in IntSlice_Uint8Slice")
	}
}
func Test_IntSlice_Uint8SlicePtr(t *testing.T) {
	var i []int
	var ii *[]uint8
	ii = IntSlice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in IntSlice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_IntSlice(t *testing.T) {
	var i []uint8
	var ii []int
	ii = Uint8Slice_IntSlice(i)
	if len(IntSlice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_IntSlice")
	}
}
func Test_Uint8Slice_IntSlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]int
	ii = Uint8Slice_IntSlicePtr(i)
	if len(IntSlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_IntSlicePtr")
	}
}
func Test_Uint8Slice_Uint8SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]uint8
	ii = Uint8Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Uint8SlicePtr")
	}
}
func Test_IntSlicePtr_Uint8Slice(t *testing.T) {
	var i *[]int
	var ii []uint8
	ii = IntSlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_IntSlice(ii)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint8Slice")
	}
}
func Test_IntSlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]int
	var ii *[]uint8
	ii = IntSlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_IntSlice(ii, UseDefaultEmpty)) != len(IntSlicePtr_IntSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in IntSlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_IntSlice(t *testing.T) {
	var i *[]uint8
	var ii []int
	ii = Uint8SlicePtr_IntSlice(i, UseDefaultEmpty)
	if len(IntSlice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_IntSlice")
	}
}
func Test_Uint8SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]uint8
	var ii []uint8
	ii = Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Uint8Slice")
	}
}
func Test_Uint8SlicePtr_IntSlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]int
	ii = Uint8SlicePtr_IntSlicePtr(i, UseDefaultEmpty)
	if len(IntSlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_IntSlicePtr")
	}
}

func Test_Int16_Int32(t *testing.T) {
	var i int16
	var ii int32
	ii = Int16_Int32(i)
	if Int32_Int16(ii) != i {
		t.Errorf("error in Int16_Int32")
	}
}
func Test_Int16_Int32Ptr(t *testing.T) {
	var i int16
	var ii *int32
	ii = Int16_Int32Ptr(i)
	if Int32Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Int32Ptr")
	}
}
func Test_Int32_Int16(t *testing.T) {
	var i int32
	var ii int16
	ii = Int32_Int16(i)
	if Int16_Int32(ii) != i {
		t.Errorf("error in Int32_Int16")
	}
}
func Test_Int32_Int16Ptr(t *testing.T) {
	var i int32
	var ii *int16
	ii = Int32_Int16Ptr(i)
	if Int16Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Int16Ptr")
	}
}
func Test_Int16Ptr_Int32(t *testing.T) {
	var i *int16
	var ii int32
	ii = Int16Ptr_Int32(i, UseDefaultEmpty)
	if Int32_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int32")
	}
}
func Test_Int16Ptr_Int32Ptr(t *testing.T) {
	var i *int16
	var ii *int32
	ii = Int16Ptr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int32Ptr")
	}
}
func Test_Int32Ptr_Int16(t *testing.T) {
	var i *int32
	var ii int16
	ii = Int32Ptr_Int16(i, UseDefaultEmpty)
	if Int16_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int16")
	}
}
func Test_Int32Ptr_Int16Ptr(t *testing.T) {
	var i *int32
	var ii *int16
	ii = Int32Ptr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int16Ptr")
	}
}

func Test_Int16Slice_Int32Slice(t *testing.T) {
	var i []int16
	var ii []int32
	ii = Int16Slice_Int32Slice(i)
	if len(Int32Slice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_Int32Slice")
	}
}
func Test_Int16Slice_Int32SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]int32
	ii = Int16Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Int32SlicePtr")
	}
}
func Test_Int32Slice_Int16Slice(t *testing.T) {
	var i []int32
	var ii []int16
	ii = Int32Slice_Int16Slice(i)
	if len(Int16Slice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_Int16Slice")
	}
}
func Test_Int32Slice_Int16SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]int16
	ii = Int32Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_Int32Slice(t *testing.T) {
	var i *[]int16
	var ii []int32
	ii = Int16SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Int32Slice")
	}
}
func Test_Int16SlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]int32
	ii = Int16SlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_Int16Slice(t *testing.T) {
	var i *[]int32
	var ii []int16
	ii = Int32SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Int16Slice")
	}
}
func Test_Int32SlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]int16
	ii = Int32SlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Int16SlicePtr")
	}
}

func Test_Int16_Int64(t *testing.T) {
	var i int16
	var ii int64
	ii = Int16_Int64(i)
	if Int64_Int16(ii) != i {
		t.Errorf("error in Int16_Int64")
	}
}
func Test_Int16_Int64Ptr(t *testing.T) {
	var i int16
	var ii *int64
	ii = Int16_Int64Ptr(i)
	if Int64Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Int64Ptr")
	}
}
func Test_Int64_Int16(t *testing.T) {
	var i int64
	var ii int16
	ii = Int64_Int16(i)
	if Int16_Int64(ii) != i {
		t.Errorf("error in Int64_Int16")
	}
}
func Test_Int64_Int16Ptr(t *testing.T) {
	var i int64
	var ii *int16
	ii = Int64_Int16Ptr(i)
	if Int16Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Int16Ptr")
	}
}
func Test_Int16Ptr_Int64(t *testing.T) {
	var i *int16
	var ii int64
	ii = Int16Ptr_Int64(i, UseDefaultEmpty)
	if Int64_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int64")
	}
}
func Test_Int16Ptr_Int64Ptr(t *testing.T) {
	var i *int16
	var ii *int64
	ii = Int16Ptr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int64Ptr")
	}
}
func Test_Int64Ptr_Int16(t *testing.T) {
	var i *int64
	var ii int16
	ii = Int64Ptr_Int16(i, UseDefaultEmpty)
	if Int16_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int16")
	}
}
func Test_Int64Ptr_Int16Ptr(t *testing.T) {
	var i *int64
	var ii *int16
	ii = Int64Ptr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int16Ptr")
	}
}

func Test_Int16Slice_Int64Slice(t *testing.T) {
	var i []int16
	var ii []int64
	ii = Int16Slice_Int64Slice(i)
	if len(Int64Slice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_Int64Slice")
	}
}
func Test_Int16Slice_Int64SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]int64
	ii = Int16Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Int64SlicePtr")
	}
}
func Test_Int64Slice_Int16Slice(t *testing.T) {
	var i []int64
	var ii []int16
	ii = Int64Slice_Int16Slice(i)
	if len(Int16Slice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_Int16Slice")
	}
}
func Test_Int64Slice_Int16SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]int16
	ii = Int64Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_Int64Slice(t *testing.T) {
	var i *[]int16
	var ii []int64
	ii = Int16SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Int64Slice")
	}
}
func Test_Int16SlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]int64
	ii = Int16SlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_Int16Slice(t *testing.T) {
	var i *[]int64
	var ii []int16
	ii = Int64SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Int16Slice")
	}
}
func Test_Int64SlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]int16
	ii = Int64SlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Int16SlicePtr")
	}
}

func Test_Int16_Int8(t *testing.T) {
	var i int16
	var ii int8
	ii = Int16_Int8(i)
	if Int8_Int16(ii) != i {
		t.Errorf("error in Int16_Int8")
	}
}
func Test_Int16_Int8Ptr(t *testing.T) {
	var i int16
	var ii *int8
	ii = Int16_Int8Ptr(i)
	if Int8Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Int8Ptr")
	}
}
func Test_Int8_Int16(t *testing.T) {
	var i int8
	var ii int16
	ii = Int8_Int16(i)
	if Int16_Int8(ii) != i {
		t.Errorf("error in Int8_Int16")
	}
}
func Test_Int8_Int16Ptr(t *testing.T) {
	var i int8
	var ii *int16
	ii = Int8_Int16Ptr(i)
	if Int16Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Int16Ptr")
	}
}
func Test_Int16Ptr_Int8(t *testing.T) {
	var i *int16
	var ii int8
	ii = Int16Ptr_Int8(i, UseDefaultEmpty)
	if Int8_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int8")
	}
}
func Test_Int16Ptr_Int8Ptr(t *testing.T) {
	var i *int16
	var ii *int8
	ii = Int16Ptr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Int8Ptr")
	}
}
func Test_Int8Ptr_Int16(t *testing.T) {
	var i *int8
	var ii int16
	ii = Int8Ptr_Int16(i, UseDefaultEmpty)
	if Int16_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int16")
	}
}
func Test_Int8Ptr_Int16Ptr(t *testing.T) {
	var i *int8
	var ii *int16
	ii = Int8Ptr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int16Ptr")
	}
}

func Test_Int16Slice_Int8Slice(t *testing.T) {
	var i []int16
	var ii []int8
	ii = Int16Slice_Int8Slice(i)
	if len(Int8Slice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_Int8Slice")
	}
}
func Test_Int16Slice_Int8SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]int8
	ii = Int16Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Int8SlicePtr")
	}
}
func Test_Int8Slice_Int16Slice(t *testing.T) {
	var i []int8
	var ii []int16
	ii = Int8Slice_Int16Slice(i)
	if len(Int16Slice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_Int16Slice")
	}
}
func Test_Int8Slice_Int16SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]int16
	ii = Int8Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_Int8Slice(t *testing.T) {
	var i *[]int16
	var ii []int8
	ii = Int16SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Int8Slice")
	}
}
func Test_Int16SlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]int8
	ii = Int16SlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_Int16Slice(t *testing.T) {
	var i *[]int8
	var ii []int16
	ii = Int8SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Int16Slice")
	}
}
func Test_Int8SlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]int16
	ii = Int8SlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Int16SlicePtr")
	}
}

func Test_Int16_Uint(t *testing.T) {
	var i int16
	var ii uint
	ii = Int16_Uint(i)
	if Uint_Int16(ii) != i {
		t.Errorf("error in Int16_Uint")
	}
}
func Test_Int16_UintPtr(t *testing.T) {
	var i int16
	var ii *uint
	ii = Int16_UintPtr(i)
	if UintPtr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_UintPtr")
	}
}
func Test_Uint_Int16(t *testing.T) {
	var i uint
	var ii int16
	ii = Uint_Int16(i)
	if Int16_Uint(ii) != i {
		t.Errorf("error in Uint_Int16")
	}
}
func Test_Uint_Int16Ptr(t *testing.T) {
	var i uint
	var ii *int16
	ii = Uint_Int16Ptr(i)
	if Int16Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Int16Ptr")
	}
}
func Test_Int16Ptr_Uint(t *testing.T) {
	var i *int16
	var ii uint
	ii = Int16Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint")
	}
}
func Test_Int16Ptr_UintPtr(t *testing.T) {
	var i *int16
	var ii *uint
	ii = Int16Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_UintPtr")
	}
}
func Test_UintPtr_Int16(t *testing.T) {
	var i *uint
	var ii int16
	ii = UintPtr_Int16(i, UseDefaultEmpty)
	if Int16_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int16")
	}
}
func Test_UintPtr_Int16Ptr(t *testing.T) {
	var i *uint
	var ii *int16
	ii = UintPtr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int16Ptr")
	}
}

func Test_Int16Slice_UintSlice(t *testing.T) {
	var i []int16
	var ii []uint
	ii = Int16Slice_UintSlice(i)
	if len(UintSlice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_UintSlice")
	}
}
func Test_Int16Slice_UintSlicePtr(t *testing.T) {
	var i []int16
	var ii *[]uint
	ii = Int16Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_UintSlicePtr")
	}
}
func Test_UintSlice_Int16Slice(t *testing.T) {
	var i []uint
	var ii []int16
	ii = UintSlice_Int16Slice(i)
	if len(Int16Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Int16Slice")
	}
}
func Test_UintSlice_Int16SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]int16
	ii = UintSlice_Int16SlicePtr(i)
	if len(Int16SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_UintSlice(t *testing.T) {
	var i *[]int16
	var ii []uint
	ii = Int16SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_UintSlice")
	}
}
func Test_Int16SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]uint
	ii = Int16SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Int16Slice(t *testing.T) {
	var i *[]uint
	var ii []int16
	ii = UintSlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int16Slice")
	}
}
func Test_UintSlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]int16
	ii = UintSlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int16SlicePtr")
	}
}

func Test_Int16_Uint16(t *testing.T) {
	var i int16
	var ii uint16
	ii = Int16_Uint16(i)
	if Uint16_Int16(ii) != i {
		t.Errorf("error in Int16_Uint16")
	}
}
func Test_Int16_Uint16Ptr(t *testing.T) {
	var i int16
	var ii *uint16
	ii = Int16_Uint16Ptr(i)
	if Uint16Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Uint16Ptr")
	}
}
func Test_Uint16_Int16(t *testing.T) {
	var i uint16
	var ii int16
	ii = Uint16_Int16(i)
	if Int16_Uint16(ii) != i {
		t.Errorf("error in Uint16_Int16")
	}
}
func Test_Uint16_Int16Ptr(t *testing.T) {
	var i uint16
	var ii *int16
	ii = Uint16_Int16Ptr(i)
	if Int16Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Int16Ptr")
	}
}
func Test_Int16Ptr_Uint16(t *testing.T) {
	var i *int16
	var ii uint16
	ii = Int16Ptr_Uint16(i, UseDefaultEmpty)
	if Uint16_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint16")
	}
}
func Test_Int16Ptr_Uint16Ptr(t *testing.T) {
	var i *int16
	var ii *uint16
	ii = Int16Ptr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Int16(t *testing.T) {
	var i *uint16
	var ii int16
	ii = Uint16Ptr_Int16(i, UseDefaultEmpty)
	if Int16_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int16")
	}
}
func Test_Uint16Ptr_Int16Ptr(t *testing.T) {
	var i *uint16
	var ii *int16
	ii = Uint16Ptr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int16Ptr")
	}
}

func Test_Int16Slice_Uint16Slice(t *testing.T) {
	var i []int16
	var ii []uint16
	ii = Int16Slice_Uint16Slice(i)
	if len(Uint16Slice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_Uint16Slice")
	}
}
func Test_Int16Slice_Uint16SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]uint16
	ii = Int16Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Uint16SlicePtr")
	}
}
func Test_Uint16Slice_Int16Slice(t *testing.T) {
	var i []uint16
	var ii []int16
	ii = Uint16Slice_Int16Slice(i)
	if len(Int16Slice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_Int16Slice")
	}
}
func Test_Uint16Slice_Int16SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]int16
	ii = Uint16Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]int16
	var ii []uint16
	ii = Int16SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint16Slice")
	}
}
func Test_Int16SlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]uint16
	ii = Int16SlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_Int16Slice(t *testing.T) {
	var i *[]uint16
	var ii []int16
	ii = Uint16SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int16Slice")
	}
}
func Test_Uint16SlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]int16
	ii = Uint16SlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int16SlicePtr")
	}
}

func Test_Int16_Uint32(t *testing.T) {
	var i int16
	var ii uint32
	ii = Int16_Uint32(i)
	if Uint32_Int16(ii) != i {
		t.Errorf("error in Int16_Uint32")
	}
}
func Test_Int16_Uint32Ptr(t *testing.T) {
	var i int16
	var ii *uint32
	ii = Int16_Uint32Ptr(i)
	if Uint32Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Uint32Ptr")
	}
}
func Test_Uint32_Int16(t *testing.T) {
	var i uint32
	var ii int16
	ii = Uint32_Int16(i)
	if Int16_Uint32(ii) != i {
		t.Errorf("error in Uint32_Int16")
	}
}
func Test_Uint32_Int16Ptr(t *testing.T) {
	var i uint32
	var ii *int16
	ii = Uint32_Int16Ptr(i)
	if Int16Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Int16Ptr")
	}
}
func Test_Int16Ptr_Uint32(t *testing.T) {
	var i *int16
	var ii uint32
	ii = Int16Ptr_Uint32(i, UseDefaultEmpty)
	if Uint32_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint32")
	}
}
func Test_Int16Ptr_Uint32Ptr(t *testing.T) {
	var i *int16
	var ii *uint32
	ii = Int16Ptr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Int16(t *testing.T) {
	var i *uint32
	var ii int16
	ii = Uint32Ptr_Int16(i, UseDefaultEmpty)
	if Int16_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int16")
	}
}
func Test_Uint32Ptr_Int16Ptr(t *testing.T) {
	var i *uint32
	var ii *int16
	ii = Uint32Ptr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int16Ptr")
	}
}

func Test_Int16Slice_Uint32Slice(t *testing.T) {
	var i []int16
	var ii []uint32
	ii = Int16Slice_Uint32Slice(i)
	if len(Uint32Slice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_Uint32Slice")
	}
}
func Test_Int16Slice_Uint32SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]uint32
	ii = Int16Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Uint32SlicePtr")
	}
}
func Test_Uint32Slice_Int16Slice(t *testing.T) {
	var i []uint32
	var ii []int16
	ii = Uint32Slice_Int16Slice(i)
	if len(Int16Slice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_Int16Slice")
	}
}
func Test_Uint32Slice_Int16SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]int16
	ii = Uint32Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]int16
	var ii []uint32
	ii = Int16SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint32Slice")
	}
}
func Test_Int16SlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]uint32
	ii = Int16SlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_Int16Slice(t *testing.T) {
	var i *[]uint32
	var ii []int16
	ii = Uint32SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int16Slice")
	}
}
func Test_Uint32SlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]int16
	ii = Uint32SlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int16SlicePtr")
	}
}

func Test_Int16_Uint64(t *testing.T) {
	var i int16
	var ii uint64
	ii = Int16_Uint64(i)
	if Uint64_Int16(ii) != i {
		t.Errorf("error in Int16_Uint64")
	}
}
func Test_Int16_Uint64Ptr(t *testing.T) {
	var i int16
	var ii *uint64
	ii = Int16_Uint64Ptr(i)
	if Uint64Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Uint64Ptr")
	}
}
func Test_Uint64_Int16(t *testing.T) {
	var i uint64
	var ii int16
	ii = Uint64_Int16(i)
	if Int16_Uint64(ii) != i {
		t.Errorf("error in Uint64_Int16")
	}
}
func Test_Uint64_Int16Ptr(t *testing.T) {
	var i uint64
	var ii *int16
	ii = Uint64_Int16Ptr(i)
	if Int16Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Int16Ptr")
	}
}
func Test_Int16Ptr_Uint64(t *testing.T) {
	var i *int16
	var ii uint64
	ii = Int16Ptr_Uint64(i, UseDefaultEmpty)
	if Uint64_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint64")
	}
}
func Test_Int16Ptr_Uint64Ptr(t *testing.T) {
	var i *int16
	var ii *uint64
	ii = Int16Ptr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Int16(t *testing.T) {
	var i *uint64
	var ii int16
	ii = Uint64Ptr_Int16(i, UseDefaultEmpty)
	if Int16_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int16")
	}
}
func Test_Uint64Ptr_Int16Ptr(t *testing.T) {
	var i *uint64
	var ii *int16
	ii = Uint64Ptr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int16Ptr")
	}
}

func Test_Int16Slice_Uint64Slice(t *testing.T) {
	var i []int16
	var ii []uint64
	ii = Int16Slice_Uint64Slice(i)
	if len(Uint64Slice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_Uint64Slice")
	}
}
func Test_Int16Slice_Uint64SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]uint64
	ii = Int16Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_Int16Slice(t *testing.T) {
	var i []uint64
	var ii []int16
	ii = Uint64Slice_Int16Slice(i)
	if len(Int16Slice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_Int16Slice")
	}
}
func Test_Uint64Slice_Int16SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]int16
	ii = Uint64Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]int16
	var ii []uint64
	ii = Int16SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint64Slice")
	}
}
func Test_Int16SlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]uint64
	ii = Int16SlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_Int16Slice(t *testing.T) {
	var i *[]uint64
	var ii []int16
	ii = Uint64SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int16Slice")
	}
}
func Test_Uint64SlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]int16
	ii = Uint64SlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int16SlicePtr")
	}
}

func Test_Int16_Uint8(t *testing.T) {
	var i int16
	var ii uint8
	ii = Int16_Uint8(i)
	if Uint8_Int16(ii) != i {
		t.Errorf("error in Int16_Uint8")
	}
}
func Test_Int16_Uint8Ptr(t *testing.T) {
	var i int16
	var ii *uint8
	ii = Int16_Uint8Ptr(i)
	if Uint8Ptr_Int16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int16_Uint8Ptr")
	}
}
func Test_Uint8_Int16(t *testing.T) {
	var i uint8
	var ii int16
	ii = Uint8_Int16(i)
	if Int16_Uint8(ii) != i {
		t.Errorf("error in Uint8_Int16")
	}
}
func Test_Uint8_Int16Ptr(t *testing.T) {
	var i uint8
	var ii *int16
	ii = Uint8_Int16Ptr(i)
	if Int16Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Int16Ptr")
	}
}
func Test_Int16Ptr_Uint8(t *testing.T) {
	var i *int16
	var ii uint8
	ii = Int16Ptr_Uint8(i, UseDefaultEmpty)
	if Uint8_Int16(ii) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint8")
	}
}
func Test_Int16Ptr_Uint8Ptr(t *testing.T) {
	var i *int16
	var ii *uint8
	ii = Int16Ptr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Int16(ii, UseDefaultEmpty) != Int16Ptr_Int16(i, UseDefaultEmpty) {
		t.Errorf("error in Int16Ptr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Int16(t *testing.T) {
	var i *uint8
	var ii int16
	ii = Uint8Ptr_Int16(i, UseDefaultEmpty)
	if Int16_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int16")
	}
}
func Test_Uint8Ptr_Int16Ptr(t *testing.T) {
	var i *uint8
	var ii *int16
	ii = Uint8Ptr_Int16Ptr(i, UseDefaultEmpty)
	if Int16Ptr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int16Ptr")
	}
}

func Test_Int16Slice_Uint8Slice(t *testing.T) {
	var i []int16
	var ii []uint8
	ii = Int16Slice_Uint8Slice(i)
	if len(Uint8Slice_Int16Slice(ii)) != len(i) {
		t.Errorf("error in Int16Slice_Uint8Slice")
	}
}
func Test_Int16Slice_Uint8SlicePtr(t *testing.T) {
	var i []int16
	var ii *[]uint8
	ii = Int16Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int16Slice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_Int16Slice(t *testing.T) {
	var i []uint8
	var ii []int16
	ii = Uint8Slice_Int16Slice(i)
	if len(Int16Slice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_Int16Slice")
	}
}
func Test_Uint8Slice_Int16SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]int16
	ii = Uint8Slice_Int16SlicePtr(i)
	if len(Int16SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Int16SlicePtr")
	}
}
func Test_Int16SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]int16
	var ii []uint8
	ii = Int16SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_Int16Slice(ii)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint8Slice")
	}
}
func Test_Int16SlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]int16
	var ii *[]uint8
	ii = Int16SlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_Int16Slice(ii, UseDefaultEmpty)) != len(Int16SlicePtr_Int16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int16SlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_Int16Slice(t *testing.T) {
	var i *[]uint8
	var ii []int16
	ii = Uint8SlicePtr_Int16Slice(i, UseDefaultEmpty)
	if len(Int16Slice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int16Slice")
	}
}
func Test_Uint8SlicePtr_Int16SlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]int16
	ii = Uint8SlicePtr_Int16SlicePtr(i, UseDefaultEmpty)
	if len(Int16SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int16SlicePtr")
	}
}

func Test_Int32_Int64(t *testing.T) {
	var i int32
	var ii int64
	ii = Int32_Int64(i)
	if Int64_Int32(ii) != i {
		t.Errorf("error in Int32_Int64")
	}
}
func Test_Int32_Int64Ptr(t *testing.T) {
	var i int32
	var ii *int64
	ii = Int32_Int64Ptr(i)
	if Int64Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Int64Ptr")
	}
}
func Test_Int64_Int32(t *testing.T) {
	var i int64
	var ii int32
	ii = Int64_Int32(i)
	if Int32_Int64(ii) != i {
		t.Errorf("error in Int64_Int32")
	}
}
func Test_Int64_Int32Ptr(t *testing.T) {
	var i int64
	var ii *int32
	ii = Int64_Int32Ptr(i)
	if Int32Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Int32Ptr")
	}
}
func Test_Int32Ptr_Int64(t *testing.T) {
	var i *int32
	var ii int64
	ii = Int32Ptr_Int64(i, UseDefaultEmpty)
	if Int64_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int64")
	}
}
func Test_Int32Ptr_Int64Ptr(t *testing.T) {
	var i *int32
	var ii *int64
	ii = Int32Ptr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int64Ptr")
	}
}
func Test_Int64Ptr_Int32(t *testing.T) {
	var i *int64
	var ii int32
	ii = Int64Ptr_Int32(i, UseDefaultEmpty)
	if Int32_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int32")
	}
}
func Test_Int64Ptr_Int32Ptr(t *testing.T) {
	var i *int64
	var ii *int32
	ii = Int64Ptr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int32Ptr")
	}
}

func Test_Int32Slice_Int64Slice(t *testing.T) {
	var i []int32
	var ii []int64
	ii = Int32Slice_Int64Slice(i)
	if len(Int64Slice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_Int64Slice")
	}
}
func Test_Int32Slice_Int64SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]int64
	ii = Int32Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Int64SlicePtr")
	}
}
func Test_Int64Slice_Int32Slice(t *testing.T) {
	var i []int64
	var ii []int32
	ii = Int64Slice_Int32Slice(i)
	if len(Int32Slice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_Int32Slice")
	}
}
func Test_Int64Slice_Int32SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]int32
	ii = Int64Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_Int64Slice(t *testing.T) {
	var i *[]int32
	var ii []int64
	ii = Int32SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Int64Slice")
	}
}
func Test_Int32SlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]int64
	ii = Int32SlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_Int32Slice(t *testing.T) {
	var i *[]int64
	var ii []int32
	ii = Int64SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Int32Slice")
	}
}
func Test_Int64SlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]int32
	ii = Int64SlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Int32SlicePtr")
	}
}

func Test_Int32_Int8(t *testing.T) {
	var i int32
	var ii int8
	ii = Int32_Int8(i)
	if Int8_Int32(ii) != i {
		t.Errorf("error in Int32_Int8")
	}
}
func Test_Int32_Int8Ptr(t *testing.T) {
	var i int32
	var ii *int8
	ii = Int32_Int8Ptr(i)
	if Int8Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Int8Ptr")
	}
}
func Test_Int8_Int32(t *testing.T) {
	var i int8
	var ii int32
	ii = Int8_Int32(i)
	if Int32_Int8(ii) != i {
		t.Errorf("error in Int8_Int32")
	}
}
func Test_Int8_Int32Ptr(t *testing.T) {
	var i int8
	var ii *int32
	ii = Int8_Int32Ptr(i)
	if Int32Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Int32Ptr")
	}
}
func Test_Int32Ptr_Int8(t *testing.T) {
	var i *int32
	var ii int8
	ii = Int32Ptr_Int8(i, UseDefaultEmpty)
	if Int8_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int8")
	}
}
func Test_Int32Ptr_Int8Ptr(t *testing.T) {
	var i *int32
	var ii *int8
	ii = Int32Ptr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Int8Ptr")
	}
}
func Test_Int8Ptr_Int32(t *testing.T) {
	var i *int8
	var ii int32
	ii = Int8Ptr_Int32(i, UseDefaultEmpty)
	if Int32_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int32")
	}
}
func Test_Int8Ptr_Int32Ptr(t *testing.T) {
	var i *int8
	var ii *int32
	ii = Int8Ptr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int32Ptr")
	}
}

func Test_Int32Slice_Int8Slice(t *testing.T) {
	var i []int32
	var ii []int8
	ii = Int32Slice_Int8Slice(i)
	if len(Int8Slice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_Int8Slice")
	}
}
func Test_Int32Slice_Int8SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]int8
	ii = Int32Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Int8SlicePtr")
	}
}
func Test_Int8Slice_Int32Slice(t *testing.T) {
	var i []int8
	var ii []int32
	ii = Int8Slice_Int32Slice(i)
	if len(Int32Slice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_Int32Slice")
	}
}
func Test_Int8Slice_Int32SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]int32
	ii = Int8Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_Int8Slice(t *testing.T) {
	var i *[]int32
	var ii []int8
	ii = Int32SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Int8Slice")
	}
}
func Test_Int32SlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]int8
	ii = Int32SlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_Int32Slice(t *testing.T) {
	var i *[]int8
	var ii []int32
	ii = Int8SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Int32Slice")
	}
}
func Test_Int8SlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]int32
	ii = Int8SlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Int32SlicePtr")
	}
}

func Test_Int32_Uint(t *testing.T) {
	var i int32
	var ii uint
	ii = Int32_Uint(i)
	if Uint_Int32(ii) != i {
		t.Errorf("error in Int32_Uint")
	}
}
func Test_Int32_UintPtr(t *testing.T) {
	var i int32
	var ii *uint
	ii = Int32_UintPtr(i)
	if UintPtr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_UintPtr")
	}
}
func Test_Uint_Int32(t *testing.T) {
	var i uint
	var ii int32
	ii = Uint_Int32(i)
	if Int32_Uint(ii) != i {
		t.Errorf("error in Uint_Int32")
	}
}
func Test_Uint_Int32Ptr(t *testing.T) {
	var i uint
	var ii *int32
	ii = Uint_Int32Ptr(i)
	if Int32Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Int32Ptr")
	}
}
func Test_Int32Ptr_Uint(t *testing.T) {
	var i *int32
	var ii uint
	ii = Int32Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint")
	}
}
func Test_Int32Ptr_UintPtr(t *testing.T) {
	var i *int32
	var ii *uint
	ii = Int32Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_UintPtr")
	}
}
func Test_UintPtr_Int32(t *testing.T) {
	var i *uint
	var ii int32
	ii = UintPtr_Int32(i, UseDefaultEmpty)
	if Int32_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int32")
	}
}
func Test_UintPtr_Int32Ptr(t *testing.T) {
	var i *uint
	var ii *int32
	ii = UintPtr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int32Ptr")
	}
}

func Test_Int32Slice_UintSlice(t *testing.T) {
	var i []int32
	var ii []uint
	ii = Int32Slice_UintSlice(i)
	if len(UintSlice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_UintSlice")
	}
}
func Test_Int32Slice_UintSlicePtr(t *testing.T) {
	var i []int32
	var ii *[]uint
	ii = Int32Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_UintSlicePtr")
	}
}
func Test_UintSlice_Int32Slice(t *testing.T) {
	var i []uint
	var ii []int32
	ii = UintSlice_Int32Slice(i)
	if len(Int32Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Int32Slice")
	}
}
func Test_UintSlice_Int32SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]int32
	ii = UintSlice_Int32SlicePtr(i)
	if len(Int32SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_UintSlice(t *testing.T) {
	var i *[]int32
	var ii []uint
	ii = Int32SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_UintSlice")
	}
}
func Test_Int32SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]uint
	ii = Int32SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Int32Slice(t *testing.T) {
	var i *[]uint
	var ii []int32
	ii = UintSlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int32Slice")
	}
}
func Test_UintSlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]int32
	ii = UintSlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int32SlicePtr")
	}
}

func Test_Int32_Uint16(t *testing.T) {
	var i int32
	var ii uint16
	ii = Int32_Uint16(i)
	if Uint16_Int32(ii) != i {
		t.Errorf("error in Int32_Uint16")
	}
}
func Test_Int32_Uint16Ptr(t *testing.T) {
	var i int32
	var ii *uint16
	ii = Int32_Uint16Ptr(i)
	if Uint16Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Uint16Ptr")
	}
}
func Test_Uint16_Int32(t *testing.T) {
	var i uint16
	var ii int32
	ii = Uint16_Int32(i)
	if Int32_Uint16(ii) != i {
		t.Errorf("error in Uint16_Int32")
	}
}
func Test_Uint16_Int32Ptr(t *testing.T) {
	var i uint16
	var ii *int32
	ii = Uint16_Int32Ptr(i)
	if Int32Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Int32Ptr")
	}
}
func Test_Int32Ptr_Uint16(t *testing.T) {
	var i *int32
	var ii uint16
	ii = Int32Ptr_Uint16(i, UseDefaultEmpty)
	if Uint16_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint16")
	}
}
func Test_Int32Ptr_Uint16Ptr(t *testing.T) {
	var i *int32
	var ii *uint16
	ii = Int32Ptr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Int32(t *testing.T) {
	var i *uint16
	var ii int32
	ii = Uint16Ptr_Int32(i, UseDefaultEmpty)
	if Int32_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int32")
	}
}
func Test_Uint16Ptr_Int32Ptr(t *testing.T) {
	var i *uint16
	var ii *int32
	ii = Uint16Ptr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int32Ptr")
	}
}

func Test_Int32Slice_Uint16Slice(t *testing.T) {
	var i []int32
	var ii []uint16
	ii = Int32Slice_Uint16Slice(i)
	if len(Uint16Slice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_Uint16Slice")
	}
}
func Test_Int32Slice_Uint16SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]uint16
	ii = Int32Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Uint16SlicePtr")
	}
}
func Test_Uint16Slice_Int32Slice(t *testing.T) {
	var i []uint16
	var ii []int32
	ii = Uint16Slice_Int32Slice(i)
	if len(Int32Slice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_Int32Slice")
	}
}
func Test_Uint16Slice_Int32SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]int32
	ii = Uint16Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]int32
	var ii []uint16
	ii = Int32SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint16Slice")
	}
}
func Test_Int32SlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]uint16
	ii = Int32SlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_Int32Slice(t *testing.T) {
	var i *[]uint16
	var ii []int32
	ii = Uint16SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int32Slice")
	}
}
func Test_Uint16SlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]int32
	ii = Uint16SlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int32SlicePtr")
	}
}

func Test_Int32_Uint32(t *testing.T) {
	var i int32
	var ii uint32
	ii = Int32_Uint32(i)
	if Uint32_Int32(ii) != i {
		t.Errorf("error in Int32_Uint32")
	}
}
func Test_Int32_Uint32Ptr(t *testing.T) {
	var i int32
	var ii *uint32
	ii = Int32_Uint32Ptr(i)
	if Uint32Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Uint32Ptr")
	}
}
func Test_Uint32_Int32(t *testing.T) {
	var i uint32
	var ii int32
	ii = Uint32_Int32(i)
	if Int32_Uint32(ii) != i {
		t.Errorf("error in Uint32_Int32")
	}
}
func Test_Uint32_Int32Ptr(t *testing.T) {
	var i uint32
	var ii *int32
	ii = Uint32_Int32Ptr(i)
	if Int32Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Int32Ptr")
	}
}
func Test_Int32Ptr_Uint32(t *testing.T) {
	var i *int32
	var ii uint32
	ii = Int32Ptr_Uint32(i, UseDefaultEmpty)
	if Uint32_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint32")
	}
}
func Test_Int32Ptr_Uint32Ptr(t *testing.T) {
	var i *int32
	var ii *uint32
	ii = Int32Ptr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Int32(t *testing.T) {
	var i *uint32
	var ii int32
	ii = Uint32Ptr_Int32(i, UseDefaultEmpty)
	if Int32_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int32")
	}
}
func Test_Uint32Ptr_Int32Ptr(t *testing.T) {
	var i *uint32
	var ii *int32
	ii = Uint32Ptr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int32Ptr")
	}
}

func Test_Int32Slice_Uint32Slice(t *testing.T) {
	var i []int32
	var ii []uint32
	ii = Int32Slice_Uint32Slice(i)
	if len(Uint32Slice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_Uint32Slice")
	}
}
func Test_Int32Slice_Uint32SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]uint32
	ii = Int32Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Uint32SlicePtr")
	}
}
func Test_Uint32Slice_Int32Slice(t *testing.T) {
	var i []uint32
	var ii []int32
	ii = Uint32Slice_Int32Slice(i)
	if len(Int32Slice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_Int32Slice")
	}
}
func Test_Uint32Slice_Int32SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]int32
	ii = Uint32Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]int32
	var ii []uint32
	ii = Int32SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint32Slice")
	}
}
func Test_Int32SlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]uint32
	ii = Int32SlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_Int32Slice(t *testing.T) {
	var i *[]uint32
	var ii []int32
	ii = Uint32SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int32Slice")
	}
}
func Test_Uint32SlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]int32
	ii = Uint32SlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int32SlicePtr")
	}
}

func Test_Int32_Uint64(t *testing.T) {
	var i int32
	var ii uint64
	ii = Int32_Uint64(i)
	if Uint64_Int32(ii) != i {
		t.Errorf("error in Int32_Uint64")
	}
}
func Test_Int32_Uint64Ptr(t *testing.T) {
	var i int32
	var ii *uint64
	ii = Int32_Uint64Ptr(i)
	if Uint64Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Uint64Ptr")
	}
}
func Test_Uint64_Int32(t *testing.T) {
	var i uint64
	var ii int32
	ii = Uint64_Int32(i)
	if Int32_Uint64(ii) != i {
		t.Errorf("error in Uint64_Int32")
	}
}
func Test_Uint64_Int32Ptr(t *testing.T) {
	var i uint64
	var ii *int32
	ii = Uint64_Int32Ptr(i)
	if Int32Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Int32Ptr")
	}
}
func Test_Int32Ptr_Uint64(t *testing.T) {
	var i *int32
	var ii uint64
	ii = Int32Ptr_Uint64(i, UseDefaultEmpty)
	if Uint64_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint64")
	}
}
func Test_Int32Ptr_Uint64Ptr(t *testing.T) {
	var i *int32
	var ii *uint64
	ii = Int32Ptr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Int32(t *testing.T) {
	var i *uint64
	var ii int32
	ii = Uint64Ptr_Int32(i, UseDefaultEmpty)
	if Int32_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int32")
	}
}
func Test_Uint64Ptr_Int32Ptr(t *testing.T) {
	var i *uint64
	var ii *int32
	ii = Uint64Ptr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int32Ptr")
	}
}

func Test_Int32Slice_Uint64Slice(t *testing.T) {
	var i []int32
	var ii []uint64
	ii = Int32Slice_Uint64Slice(i)
	if len(Uint64Slice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_Uint64Slice")
	}
}
func Test_Int32Slice_Uint64SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]uint64
	ii = Int32Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_Int32Slice(t *testing.T) {
	var i []uint64
	var ii []int32
	ii = Uint64Slice_Int32Slice(i)
	if len(Int32Slice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_Int32Slice")
	}
}
func Test_Uint64Slice_Int32SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]int32
	ii = Uint64Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]int32
	var ii []uint64
	ii = Int32SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint64Slice")
	}
}
func Test_Int32SlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]uint64
	ii = Int32SlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_Int32Slice(t *testing.T) {
	var i *[]uint64
	var ii []int32
	ii = Uint64SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int32Slice")
	}
}
func Test_Uint64SlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]int32
	ii = Uint64SlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int32SlicePtr")
	}
}

func Test_Int32_Uint8(t *testing.T) {
	var i int32
	var ii uint8
	ii = Int32_Uint8(i)
	if Uint8_Int32(ii) != i {
		t.Errorf("error in Int32_Uint8")
	}
}
func Test_Int32_Uint8Ptr(t *testing.T) {
	var i int32
	var ii *uint8
	ii = Int32_Uint8Ptr(i)
	if Uint8Ptr_Int32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int32_Uint8Ptr")
	}
}
func Test_Uint8_Int32(t *testing.T) {
	var i uint8
	var ii int32
	ii = Uint8_Int32(i)
	if Int32_Uint8(ii) != i {
		t.Errorf("error in Uint8_Int32")
	}
}
func Test_Uint8_Int32Ptr(t *testing.T) {
	var i uint8
	var ii *int32
	ii = Uint8_Int32Ptr(i)
	if Int32Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Int32Ptr")
	}
}
func Test_Int32Ptr_Uint8(t *testing.T) {
	var i *int32
	var ii uint8
	ii = Int32Ptr_Uint8(i, UseDefaultEmpty)
	if Uint8_Int32(ii) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint8")
	}
}
func Test_Int32Ptr_Uint8Ptr(t *testing.T) {
	var i *int32
	var ii *uint8
	ii = Int32Ptr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Int32(ii, UseDefaultEmpty) != Int32Ptr_Int32(i, UseDefaultEmpty) {
		t.Errorf("error in Int32Ptr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Int32(t *testing.T) {
	var i *uint8
	var ii int32
	ii = Uint8Ptr_Int32(i, UseDefaultEmpty)
	if Int32_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int32")
	}
}
func Test_Uint8Ptr_Int32Ptr(t *testing.T) {
	var i *uint8
	var ii *int32
	ii = Uint8Ptr_Int32Ptr(i, UseDefaultEmpty)
	if Int32Ptr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int32Ptr")
	}
}

func Test_Int32Slice_Uint8Slice(t *testing.T) {
	var i []int32
	var ii []uint8
	ii = Int32Slice_Uint8Slice(i)
	if len(Uint8Slice_Int32Slice(ii)) != len(i) {
		t.Errorf("error in Int32Slice_Uint8Slice")
	}
}
func Test_Int32Slice_Uint8SlicePtr(t *testing.T) {
	var i []int32
	var ii *[]uint8
	ii = Int32Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int32Slice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_Int32Slice(t *testing.T) {
	var i []uint8
	var ii []int32
	ii = Uint8Slice_Int32Slice(i)
	if len(Int32Slice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_Int32Slice")
	}
}
func Test_Uint8Slice_Int32SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]int32
	ii = Uint8Slice_Int32SlicePtr(i)
	if len(Int32SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Int32SlicePtr")
	}
}
func Test_Int32SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]int32
	var ii []uint8
	ii = Int32SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_Int32Slice(ii)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint8Slice")
	}
}
func Test_Int32SlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]int32
	var ii *[]uint8
	ii = Int32SlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_Int32Slice(ii, UseDefaultEmpty)) != len(Int32SlicePtr_Int32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int32SlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_Int32Slice(t *testing.T) {
	var i *[]uint8
	var ii []int32
	ii = Uint8SlicePtr_Int32Slice(i, UseDefaultEmpty)
	if len(Int32Slice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int32Slice")
	}
}
func Test_Uint8SlicePtr_Int32SlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]int32
	ii = Uint8SlicePtr_Int32SlicePtr(i, UseDefaultEmpty)
	if len(Int32SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int32SlicePtr")
	}
}

func Test_Int64_Int8(t *testing.T) {
	var i int64
	var ii int8
	ii = Int64_Int8(i)
	if Int8_Int64(ii) != i {
		t.Errorf("error in Int64_Int8")
	}
}
func Test_Int64_Int8Ptr(t *testing.T) {
	var i int64
	var ii *int8
	ii = Int64_Int8Ptr(i)
	if Int8Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Int8Ptr")
	}
}
func Test_Int8_Int64(t *testing.T) {
	var i int8
	var ii int64
	ii = Int8_Int64(i)
	if Int64_Int8(ii) != i {
		t.Errorf("error in Int8_Int64")
	}
}
func Test_Int8_Int64Ptr(t *testing.T) {
	var i int8
	var ii *int64
	ii = Int8_Int64Ptr(i)
	if Int64Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Int64Ptr")
	}
}
func Test_Int64Ptr_Int8(t *testing.T) {
	var i *int64
	var ii int8
	ii = Int64Ptr_Int8(i, UseDefaultEmpty)
	if Int8_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int8")
	}
}
func Test_Int64Ptr_Int8Ptr(t *testing.T) {
	var i *int64
	var ii *int8
	ii = Int64Ptr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Int8Ptr")
	}
}
func Test_Int8Ptr_Int64(t *testing.T) {
	var i *int8
	var ii int64
	ii = Int8Ptr_Int64(i, UseDefaultEmpty)
	if Int64_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int64")
	}
}
func Test_Int8Ptr_Int64Ptr(t *testing.T) {
	var i *int8
	var ii *int64
	ii = Int8Ptr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Int64Ptr")
	}
}

func Test_Int64Slice_Int8Slice(t *testing.T) {
	var i []int64
	var ii []int8
	ii = Int64Slice_Int8Slice(i)
	if len(Int8Slice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_Int8Slice")
	}
}
func Test_Int64Slice_Int8SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]int8
	ii = Int64Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Int8SlicePtr")
	}
}
func Test_Int8Slice_Int64Slice(t *testing.T) {
	var i []int8
	var ii []int64
	ii = Int8Slice_Int64Slice(i)
	if len(Int64Slice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_Int64Slice")
	}
}
func Test_Int8Slice_Int64SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]int64
	ii = Int8Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_Int8Slice(t *testing.T) {
	var i *[]int64
	var ii []int8
	ii = Int64SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Int8Slice")
	}
}
func Test_Int64SlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]int8
	ii = Int64SlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_Int64Slice(t *testing.T) {
	var i *[]int8
	var ii []int64
	ii = Int8SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Int64Slice")
	}
}
func Test_Int8SlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]int64
	ii = Int8SlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Int64SlicePtr")
	}
}

func Test_Int64_Uint(t *testing.T) {
	var i int64
	var ii uint
	ii = Int64_Uint(i)
	if Uint_Int64(ii) != i {
		t.Errorf("error in Int64_Uint")
	}
}
func Test_Int64_UintPtr(t *testing.T) {
	var i int64
	var ii *uint
	ii = Int64_UintPtr(i)
	if UintPtr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_UintPtr")
	}
}
func Test_Uint_Int64(t *testing.T) {
	var i uint
	var ii int64
	ii = Uint_Int64(i)
	if Int64_Uint(ii) != i {
		t.Errorf("error in Uint_Int64")
	}
}
func Test_Uint_Int64Ptr(t *testing.T) {
	var i uint
	var ii *int64
	ii = Uint_Int64Ptr(i)
	if Int64Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Int64Ptr")
	}
}
func Test_Int64Ptr_Uint(t *testing.T) {
	var i *int64
	var ii uint
	ii = Int64Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint")
	}
}
func Test_Int64Ptr_UintPtr(t *testing.T) {
	var i *int64
	var ii *uint
	ii = Int64Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_UintPtr")
	}
}
func Test_UintPtr_Int64(t *testing.T) {
	var i *uint
	var ii int64
	ii = UintPtr_Int64(i, UseDefaultEmpty)
	if Int64_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int64")
	}
}
func Test_UintPtr_Int64Ptr(t *testing.T) {
	var i *uint
	var ii *int64
	ii = UintPtr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int64Ptr")
	}
}

func Test_Int64Slice_UintSlice(t *testing.T) {
	var i []int64
	var ii []uint
	ii = Int64Slice_UintSlice(i)
	if len(UintSlice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_UintSlice")
	}
}
func Test_Int64Slice_UintSlicePtr(t *testing.T) {
	var i []int64
	var ii *[]uint
	ii = Int64Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_UintSlicePtr")
	}
}
func Test_UintSlice_Int64Slice(t *testing.T) {
	var i []uint
	var ii []int64
	ii = UintSlice_Int64Slice(i)
	if len(Int64Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Int64Slice")
	}
}
func Test_UintSlice_Int64SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]int64
	ii = UintSlice_Int64SlicePtr(i)
	if len(Int64SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_UintSlice(t *testing.T) {
	var i *[]int64
	var ii []uint
	ii = Int64SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_UintSlice")
	}
}
func Test_Int64SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]uint
	ii = Int64SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Int64Slice(t *testing.T) {
	var i *[]uint
	var ii []int64
	ii = UintSlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int64Slice")
	}
}
func Test_UintSlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]int64
	ii = UintSlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int64SlicePtr")
	}
}

func Test_Int64_Uint16(t *testing.T) {
	var i int64
	var ii uint16
	ii = Int64_Uint16(i)
	if Uint16_Int64(ii) != i {
		t.Errorf("error in Int64_Uint16")
	}
}
func Test_Int64_Uint16Ptr(t *testing.T) {
	var i int64
	var ii *uint16
	ii = Int64_Uint16Ptr(i)
	if Uint16Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Uint16Ptr")
	}
}
func Test_Uint16_Int64(t *testing.T) {
	var i uint16
	var ii int64
	ii = Uint16_Int64(i)
	if Int64_Uint16(ii) != i {
		t.Errorf("error in Uint16_Int64")
	}
}
func Test_Uint16_Int64Ptr(t *testing.T) {
	var i uint16
	var ii *int64
	ii = Uint16_Int64Ptr(i)
	if Int64Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Int64Ptr")
	}
}
func Test_Int64Ptr_Uint16(t *testing.T) {
	var i *int64
	var ii uint16
	ii = Int64Ptr_Uint16(i, UseDefaultEmpty)
	if Uint16_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint16")
	}
}
func Test_Int64Ptr_Uint16Ptr(t *testing.T) {
	var i *int64
	var ii *uint16
	ii = Int64Ptr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Int64(t *testing.T) {
	var i *uint16
	var ii int64
	ii = Uint16Ptr_Int64(i, UseDefaultEmpty)
	if Int64_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int64")
	}
}
func Test_Uint16Ptr_Int64Ptr(t *testing.T) {
	var i *uint16
	var ii *int64
	ii = Uint16Ptr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int64Ptr")
	}
}

func Test_Int64Slice_Uint16Slice(t *testing.T) {
	var i []int64
	var ii []uint16
	ii = Int64Slice_Uint16Slice(i)
	if len(Uint16Slice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_Uint16Slice")
	}
}
func Test_Int64Slice_Uint16SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]uint16
	ii = Int64Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Uint16SlicePtr")
	}
}
func Test_Uint16Slice_Int64Slice(t *testing.T) {
	var i []uint16
	var ii []int64
	ii = Uint16Slice_Int64Slice(i)
	if len(Int64Slice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_Int64Slice")
	}
}
func Test_Uint16Slice_Int64SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]int64
	ii = Uint16Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]int64
	var ii []uint16
	ii = Int64SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint16Slice")
	}
}
func Test_Int64SlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]uint16
	ii = Int64SlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_Int64Slice(t *testing.T) {
	var i *[]uint16
	var ii []int64
	ii = Uint16SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int64Slice")
	}
}
func Test_Uint16SlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]int64
	ii = Uint16SlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int64SlicePtr")
	}
}

func Test_Int64_Uint32(t *testing.T) {
	var i int64
	var ii uint32
	ii = Int64_Uint32(i)
	if Uint32_Int64(ii) != i {
		t.Errorf("error in Int64_Uint32")
	}
}
func Test_Int64_Uint32Ptr(t *testing.T) {
	var i int64
	var ii *uint32
	ii = Int64_Uint32Ptr(i)
	if Uint32Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Uint32Ptr")
	}
}
func Test_Uint32_Int64(t *testing.T) {
	var i uint32
	var ii int64
	ii = Uint32_Int64(i)
	if Int64_Uint32(ii) != i {
		t.Errorf("error in Uint32_Int64")
	}
}
func Test_Uint32_Int64Ptr(t *testing.T) {
	var i uint32
	var ii *int64
	ii = Uint32_Int64Ptr(i)
	if Int64Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Int64Ptr")
	}
}
func Test_Int64Ptr_Uint32(t *testing.T) {
	var i *int64
	var ii uint32
	ii = Int64Ptr_Uint32(i, UseDefaultEmpty)
	if Uint32_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint32")
	}
}
func Test_Int64Ptr_Uint32Ptr(t *testing.T) {
	var i *int64
	var ii *uint32
	ii = Int64Ptr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Int64(t *testing.T) {
	var i *uint32
	var ii int64
	ii = Uint32Ptr_Int64(i, UseDefaultEmpty)
	if Int64_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int64")
	}
}
func Test_Uint32Ptr_Int64Ptr(t *testing.T) {
	var i *uint32
	var ii *int64
	ii = Uint32Ptr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int64Ptr")
	}
}

func Test_Int64Slice_Uint32Slice(t *testing.T) {
	var i []int64
	var ii []uint32
	ii = Int64Slice_Uint32Slice(i)
	if len(Uint32Slice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_Uint32Slice")
	}
}
func Test_Int64Slice_Uint32SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]uint32
	ii = Int64Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Uint32SlicePtr")
	}
}
func Test_Uint32Slice_Int64Slice(t *testing.T) {
	var i []uint32
	var ii []int64
	ii = Uint32Slice_Int64Slice(i)
	if len(Int64Slice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_Int64Slice")
	}
}
func Test_Uint32Slice_Int64SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]int64
	ii = Uint32Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]int64
	var ii []uint32
	ii = Int64SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint32Slice")
	}
}
func Test_Int64SlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]uint32
	ii = Int64SlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_Int64Slice(t *testing.T) {
	var i *[]uint32
	var ii []int64
	ii = Uint32SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int64Slice")
	}
}
func Test_Uint32SlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]int64
	ii = Uint32SlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int64SlicePtr")
	}
}

func Test_Int64_Uint64(t *testing.T) {
	var i int64
	var ii uint64
	ii = Int64_Uint64(i)
	if Uint64_Int64(ii) != i {
		t.Errorf("error in Int64_Uint64")
	}
}
func Test_Int64_Uint64Ptr(t *testing.T) {
	var i int64
	var ii *uint64
	ii = Int64_Uint64Ptr(i)
	if Uint64Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Uint64Ptr")
	}
}
func Test_Uint64_Int64(t *testing.T) {
	var i uint64
	var ii int64
	ii = Uint64_Int64(i)
	if Int64_Uint64(ii) != i {
		t.Errorf("error in Uint64_Int64")
	}
}
func Test_Uint64_Int64Ptr(t *testing.T) {
	var i uint64
	var ii *int64
	ii = Uint64_Int64Ptr(i)
	if Int64Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Int64Ptr")
	}
}
func Test_Int64Ptr_Uint64(t *testing.T) {
	var i *int64
	var ii uint64
	ii = Int64Ptr_Uint64(i, UseDefaultEmpty)
	if Uint64_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint64")
	}
}
func Test_Int64Ptr_Uint64Ptr(t *testing.T) {
	var i *int64
	var ii *uint64
	ii = Int64Ptr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Int64(t *testing.T) {
	var i *uint64
	var ii int64
	ii = Uint64Ptr_Int64(i, UseDefaultEmpty)
	if Int64_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int64")
	}
}
func Test_Uint64Ptr_Int64Ptr(t *testing.T) {
	var i *uint64
	var ii *int64
	ii = Uint64Ptr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int64Ptr")
	}
}

func Test_Int64Slice_Uint64Slice(t *testing.T) {
	var i []int64
	var ii []uint64
	ii = Int64Slice_Uint64Slice(i)
	if len(Uint64Slice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_Uint64Slice")
	}
}
func Test_Int64Slice_Uint64SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]uint64
	ii = Int64Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_Int64Slice(t *testing.T) {
	var i []uint64
	var ii []int64
	ii = Uint64Slice_Int64Slice(i)
	if len(Int64Slice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_Int64Slice")
	}
}
func Test_Uint64Slice_Int64SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]int64
	ii = Uint64Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]int64
	var ii []uint64
	ii = Int64SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint64Slice")
	}
}
func Test_Int64SlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]uint64
	ii = Int64SlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_Int64Slice(t *testing.T) {
	var i *[]uint64
	var ii []int64
	ii = Uint64SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int64Slice")
	}
}
func Test_Uint64SlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]int64
	ii = Uint64SlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int64SlicePtr")
	}
}

func Test_Int64_Uint8(t *testing.T) {
	var i int64
	var ii uint8
	ii = Int64_Uint8(i)
	if Uint8_Int64(ii) != i {
		t.Errorf("error in Int64_Uint8")
	}
}
func Test_Int64_Uint8Ptr(t *testing.T) {
	var i int64
	var ii *uint8
	ii = Int64_Uint8Ptr(i)
	if Uint8Ptr_Int64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int64_Uint8Ptr")
	}
}
func Test_Uint8_Int64(t *testing.T) {
	var i uint8
	var ii int64
	ii = Uint8_Int64(i)
	if Int64_Uint8(ii) != i {
		t.Errorf("error in Uint8_Int64")
	}
}
func Test_Uint8_Int64Ptr(t *testing.T) {
	var i uint8
	var ii *int64
	ii = Uint8_Int64Ptr(i)
	if Int64Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Int64Ptr")
	}
}
func Test_Int64Ptr_Uint8(t *testing.T) {
	var i *int64
	var ii uint8
	ii = Int64Ptr_Uint8(i, UseDefaultEmpty)
	if Uint8_Int64(ii) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint8")
	}
}
func Test_Int64Ptr_Uint8Ptr(t *testing.T) {
	var i *int64
	var ii *uint8
	ii = Int64Ptr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Int64(ii, UseDefaultEmpty) != Int64Ptr_Int64(i, UseDefaultEmpty) {
		t.Errorf("error in Int64Ptr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Int64(t *testing.T) {
	var i *uint8
	var ii int64
	ii = Uint8Ptr_Int64(i, UseDefaultEmpty)
	if Int64_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int64")
	}
}
func Test_Uint8Ptr_Int64Ptr(t *testing.T) {
	var i *uint8
	var ii *int64
	ii = Uint8Ptr_Int64Ptr(i, UseDefaultEmpty)
	if Int64Ptr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int64Ptr")
	}
}

func Test_Int64Slice_Uint8Slice(t *testing.T) {
	var i []int64
	var ii []uint8
	ii = Int64Slice_Uint8Slice(i)
	if len(Uint8Slice_Int64Slice(ii)) != len(i) {
		t.Errorf("error in Int64Slice_Uint8Slice")
	}
}
func Test_Int64Slice_Uint8SlicePtr(t *testing.T) {
	var i []int64
	var ii *[]uint8
	ii = Int64Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int64Slice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_Int64Slice(t *testing.T) {
	var i []uint8
	var ii []int64
	ii = Uint8Slice_Int64Slice(i)
	if len(Int64Slice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_Int64Slice")
	}
}
func Test_Uint8Slice_Int64SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]int64
	ii = Uint8Slice_Int64SlicePtr(i)
	if len(Int64SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Int64SlicePtr")
	}
}
func Test_Int64SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]int64
	var ii []uint8
	ii = Int64SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_Int64Slice(ii)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint8Slice")
	}
}
func Test_Int64SlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]int64
	var ii *[]uint8
	ii = Int64SlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_Int64Slice(ii, UseDefaultEmpty)) != len(Int64SlicePtr_Int64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int64SlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_Int64Slice(t *testing.T) {
	var i *[]uint8
	var ii []int64
	ii = Uint8SlicePtr_Int64Slice(i, UseDefaultEmpty)
	if len(Int64Slice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int64Slice")
	}
}
func Test_Uint8SlicePtr_Int64SlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]int64
	ii = Uint8SlicePtr_Int64SlicePtr(i, UseDefaultEmpty)
	if len(Int64SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int64SlicePtr")
	}
}

func Test_Int8_Uint(t *testing.T) {
	var i int8
	var ii uint
	ii = Int8_Uint(i)
	if Uint_Int8(ii) != i {
		t.Errorf("error in Int8_Uint")
	}
}
func Test_Int8_UintPtr(t *testing.T) {
	var i int8
	var ii *uint
	ii = Int8_UintPtr(i)
	if UintPtr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_UintPtr")
	}
}
func Test_Uint_Int8(t *testing.T) {
	var i uint
	var ii int8
	ii = Uint_Int8(i)
	if Int8_Uint(ii) != i {
		t.Errorf("error in Uint_Int8")
	}
}
func Test_Uint_Int8Ptr(t *testing.T) {
	var i uint
	var ii *int8
	ii = Uint_Int8Ptr(i)
	if Int8Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Int8Ptr")
	}
}
func Test_Int8Ptr_Uint(t *testing.T) {
	var i *int8
	var ii uint
	ii = Int8Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint")
	}
}
func Test_Int8Ptr_UintPtr(t *testing.T) {
	var i *int8
	var ii *uint
	ii = Int8Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_UintPtr")
	}
}
func Test_UintPtr_Int8(t *testing.T) {
	var i *uint
	var ii int8
	ii = UintPtr_Int8(i, UseDefaultEmpty)
	if Int8_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int8")
	}
}
func Test_UintPtr_Int8Ptr(t *testing.T) {
	var i *uint
	var ii *int8
	ii = UintPtr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Int8Ptr")
	}
}

func Test_Int8Slice_UintSlice(t *testing.T) {
	var i []int8
	var ii []uint
	ii = Int8Slice_UintSlice(i)
	if len(UintSlice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_UintSlice")
	}
}
func Test_Int8Slice_UintSlicePtr(t *testing.T) {
	var i []int8
	var ii *[]uint
	ii = Int8Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_UintSlicePtr")
	}
}
func Test_UintSlice_Int8Slice(t *testing.T) {
	var i []uint
	var ii []int8
	ii = UintSlice_Int8Slice(i)
	if len(Int8Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Int8Slice")
	}
}
func Test_UintSlice_Int8SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]int8
	ii = UintSlice_Int8SlicePtr(i)
	if len(Int8SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_UintSlice(t *testing.T) {
	var i *[]int8
	var ii []uint
	ii = Int8SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_UintSlice")
	}
}
func Test_Int8SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]uint
	ii = Int8SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Int8Slice(t *testing.T) {
	var i *[]uint
	var ii []int8
	ii = UintSlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int8Slice")
	}
}
func Test_UintSlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]int8
	ii = UintSlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Int8SlicePtr")
	}
}

func Test_Int8_Uint16(t *testing.T) {
	var i int8
	var ii uint16
	ii = Int8_Uint16(i)
	if Uint16_Int8(ii) != i {
		t.Errorf("error in Int8_Uint16")
	}
}
func Test_Int8_Uint16Ptr(t *testing.T) {
	var i int8
	var ii *uint16
	ii = Int8_Uint16Ptr(i)
	if Uint16Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Uint16Ptr")
	}
}
func Test_Uint16_Int8(t *testing.T) {
	var i uint16
	var ii int8
	ii = Uint16_Int8(i)
	if Int8_Uint16(ii) != i {
		t.Errorf("error in Uint16_Int8")
	}
}
func Test_Uint16_Int8Ptr(t *testing.T) {
	var i uint16
	var ii *int8
	ii = Uint16_Int8Ptr(i)
	if Int8Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Int8Ptr")
	}
}
func Test_Int8Ptr_Uint16(t *testing.T) {
	var i *int8
	var ii uint16
	ii = Int8Ptr_Uint16(i, UseDefaultEmpty)
	if Uint16_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint16")
	}
}
func Test_Int8Ptr_Uint16Ptr(t *testing.T) {
	var i *int8
	var ii *uint16
	ii = Int8Ptr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Int8(t *testing.T) {
	var i *uint16
	var ii int8
	ii = Uint16Ptr_Int8(i, UseDefaultEmpty)
	if Int8_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int8")
	}
}
func Test_Uint16Ptr_Int8Ptr(t *testing.T) {
	var i *uint16
	var ii *int8
	ii = Uint16Ptr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Int8Ptr")
	}
}

func Test_Int8Slice_Uint16Slice(t *testing.T) {
	var i []int8
	var ii []uint16
	ii = Int8Slice_Uint16Slice(i)
	if len(Uint16Slice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_Uint16Slice")
	}
}
func Test_Int8Slice_Uint16SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]uint16
	ii = Int8Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Uint16SlicePtr")
	}
}
func Test_Uint16Slice_Int8Slice(t *testing.T) {
	var i []uint16
	var ii []int8
	ii = Uint16Slice_Int8Slice(i)
	if len(Int8Slice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_Int8Slice")
	}
}
func Test_Uint16Slice_Int8SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]int8
	ii = Uint16Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]int8
	var ii []uint16
	ii = Int8SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint16Slice")
	}
}
func Test_Int8SlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]uint16
	ii = Int8SlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_Int8Slice(t *testing.T) {
	var i *[]uint16
	var ii []int8
	ii = Uint16SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int8Slice")
	}
}
func Test_Uint16SlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]int8
	ii = Uint16SlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Int8SlicePtr")
	}
}

func Test_Int8_Uint32(t *testing.T) {
	var i int8
	var ii uint32
	ii = Int8_Uint32(i)
	if Uint32_Int8(ii) != i {
		t.Errorf("error in Int8_Uint32")
	}
}
func Test_Int8_Uint32Ptr(t *testing.T) {
	var i int8
	var ii *uint32
	ii = Int8_Uint32Ptr(i)
	if Uint32Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Uint32Ptr")
	}
}
func Test_Uint32_Int8(t *testing.T) {
	var i uint32
	var ii int8
	ii = Uint32_Int8(i)
	if Int8_Uint32(ii) != i {
		t.Errorf("error in Uint32_Int8")
	}
}
func Test_Uint32_Int8Ptr(t *testing.T) {
	var i uint32
	var ii *int8
	ii = Uint32_Int8Ptr(i)
	if Int8Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Int8Ptr")
	}
}
func Test_Int8Ptr_Uint32(t *testing.T) {
	var i *int8
	var ii uint32
	ii = Int8Ptr_Uint32(i, UseDefaultEmpty)
	if Uint32_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint32")
	}
}
func Test_Int8Ptr_Uint32Ptr(t *testing.T) {
	var i *int8
	var ii *uint32
	ii = Int8Ptr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Int8(t *testing.T) {
	var i *uint32
	var ii int8
	ii = Uint32Ptr_Int8(i, UseDefaultEmpty)
	if Int8_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int8")
	}
}
func Test_Uint32Ptr_Int8Ptr(t *testing.T) {
	var i *uint32
	var ii *int8
	ii = Uint32Ptr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Int8Ptr")
	}
}

func Test_Int8Slice_Uint32Slice(t *testing.T) {
	var i []int8
	var ii []uint32
	ii = Int8Slice_Uint32Slice(i)
	if len(Uint32Slice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_Uint32Slice")
	}
}
func Test_Int8Slice_Uint32SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]uint32
	ii = Int8Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Uint32SlicePtr")
	}
}
func Test_Uint32Slice_Int8Slice(t *testing.T) {
	var i []uint32
	var ii []int8
	ii = Uint32Slice_Int8Slice(i)
	if len(Int8Slice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_Int8Slice")
	}
}
func Test_Uint32Slice_Int8SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]int8
	ii = Uint32Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]int8
	var ii []uint32
	ii = Int8SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint32Slice")
	}
}
func Test_Int8SlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]uint32
	ii = Int8SlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_Int8Slice(t *testing.T) {
	var i *[]uint32
	var ii []int8
	ii = Uint32SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int8Slice")
	}
}
func Test_Uint32SlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]int8
	ii = Uint32SlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Int8SlicePtr")
	}
}

func Test_Int8_Uint64(t *testing.T) {
	var i int8
	var ii uint64
	ii = Int8_Uint64(i)
	if Uint64_Int8(ii) != i {
		t.Errorf("error in Int8_Uint64")
	}
}
func Test_Int8_Uint64Ptr(t *testing.T) {
	var i int8
	var ii *uint64
	ii = Int8_Uint64Ptr(i)
	if Uint64Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Uint64Ptr")
	}
}
func Test_Uint64_Int8(t *testing.T) {
	var i uint64
	var ii int8
	ii = Uint64_Int8(i)
	if Int8_Uint64(ii) != i {
		t.Errorf("error in Uint64_Int8")
	}
}
func Test_Uint64_Int8Ptr(t *testing.T) {
	var i uint64
	var ii *int8
	ii = Uint64_Int8Ptr(i)
	if Int8Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Int8Ptr")
	}
}
func Test_Int8Ptr_Uint64(t *testing.T) {
	var i *int8
	var ii uint64
	ii = Int8Ptr_Uint64(i, UseDefaultEmpty)
	if Uint64_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint64")
	}
}
func Test_Int8Ptr_Uint64Ptr(t *testing.T) {
	var i *int8
	var ii *uint64
	ii = Int8Ptr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Int8(t *testing.T) {
	var i *uint64
	var ii int8
	ii = Uint64Ptr_Int8(i, UseDefaultEmpty)
	if Int8_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int8")
	}
}
func Test_Uint64Ptr_Int8Ptr(t *testing.T) {
	var i *uint64
	var ii *int8
	ii = Uint64Ptr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Int8Ptr")
	}
}

func Test_Int8Slice_Uint64Slice(t *testing.T) {
	var i []int8
	var ii []uint64
	ii = Int8Slice_Uint64Slice(i)
	if len(Uint64Slice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_Uint64Slice")
	}
}
func Test_Int8Slice_Uint64SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]uint64
	ii = Int8Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_Int8Slice(t *testing.T) {
	var i []uint64
	var ii []int8
	ii = Uint64Slice_Int8Slice(i)
	if len(Int8Slice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_Int8Slice")
	}
}
func Test_Uint64Slice_Int8SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]int8
	ii = Uint64Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]int8
	var ii []uint64
	ii = Int8SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint64Slice")
	}
}
func Test_Int8SlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]uint64
	ii = Int8SlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_Int8Slice(t *testing.T) {
	var i *[]uint64
	var ii []int8
	ii = Uint64SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int8Slice")
	}
}
func Test_Uint64SlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]int8
	ii = Uint64SlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Int8SlicePtr")
	}
}

func Test_Int8_Uint8(t *testing.T) {
	var i int8
	var ii uint8
	ii = Int8_Uint8(i)
	if Uint8_Int8(ii) != i {
		t.Errorf("error in Int8_Uint8")
	}
}
func Test_Int8_Uint8Ptr(t *testing.T) {
	var i int8
	var ii *uint8
	ii = Int8_Uint8Ptr(i)
	if Uint8Ptr_Int8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Int8_Uint8Ptr")
	}
}
func Test_Uint8_Int8(t *testing.T) {
	var i uint8
	var ii int8
	ii = Uint8_Int8(i)
	if Int8_Uint8(ii) != i {
		t.Errorf("error in Uint8_Int8")
	}
}
func Test_Uint8_Int8Ptr(t *testing.T) {
	var i uint8
	var ii *int8
	ii = Uint8_Int8Ptr(i)
	if Int8Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Int8Ptr")
	}
}
func Test_Int8Ptr_Uint8(t *testing.T) {
	var i *int8
	var ii uint8
	ii = Int8Ptr_Uint8(i, UseDefaultEmpty)
	if Uint8_Int8(ii) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint8")
	}
}
func Test_Int8Ptr_Uint8Ptr(t *testing.T) {
	var i *int8
	var ii *uint8
	ii = Int8Ptr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Int8(ii, UseDefaultEmpty) != Int8Ptr_Int8(i, UseDefaultEmpty) {
		t.Errorf("error in Int8Ptr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Int8(t *testing.T) {
	var i *uint8
	var ii int8
	ii = Uint8Ptr_Int8(i, UseDefaultEmpty)
	if Int8_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int8")
	}
}
func Test_Uint8Ptr_Int8Ptr(t *testing.T) {
	var i *uint8
	var ii *int8
	ii = Uint8Ptr_Int8Ptr(i, UseDefaultEmpty)
	if Int8Ptr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Int8Ptr")
	}
}

func Test_Int8Slice_Uint8Slice(t *testing.T) {
	var i []int8
	var ii []uint8
	ii = Int8Slice_Uint8Slice(i)
	if len(Uint8Slice_Int8Slice(ii)) != len(i) {
		t.Errorf("error in Int8Slice_Uint8Slice")
	}
}
func Test_Int8Slice_Uint8SlicePtr(t *testing.T) {
	var i []int8
	var ii *[]uint8
	ii = Int8Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Int8Slice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_Int8Slice(t *testing.T) {
	var i []uint8
	var ii []int8
	ii = Uint8Slice_Int8Slice(i)
	if len(Int8Slice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_Int8Slice")
	}
}
func Test_Uint8Slice_Int8SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]int8
	ii = Uint8Slice_Int8SlicePtr(i)
	if len(Int8SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Int8SlicePtr")
	}
}
func Test_Int8SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]int8
	var ii []uint8
	ii = Int8SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_Int8Slice(ii)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint8Slice")
	}
}
func Test_Int8SlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]int8
	var ii *[]uint8
	ii = Int8SlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_Int8Slice(ii, UseDefaultEmpty)) != len(Int8SlicePtr_Int8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Int8SlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_Int8Slice(t *testing.T) {
	var i *[]uint8
	var ii []int8
	ii = Uint8SlicePtr_Int8Slice(i, UseDefaultEmpty)
	if len(Int8Slice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int8Slice")
	}
}
func Test_Uint8SlicePtr_Int8SlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]int8
	ii = Uint8SlicePtr_Int8SlicePtr(i, UseDefaultEmpty)
	if len(Int8SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Int8SlicePtr")
	}
}

func Test_Uint_Uint16(t *testing.T) {
	var i uint
	var ii uint16
	ii = Uint_Uint16(i)
	if Uint16_Uint(ii) != i {
		t.Errorf("error in Uint_Uint16")
	}
}
func Test_Uint_Uint16Ptr(t *testing.T) {
	var i uint
	var ii *uint16
	ii = Uint_Uint16Ptr(i)
	if Uint16Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Uint16Ptr")
	}
}
func Test_Uint16_Uint(t *testing.T) {
	var i uint16
	var ii uint
	ii = Uint16_Uint(i)
	if Uint_Uint16(ii) != i {
		t.Errorf("error in Uint16_Uint")
	}
}
func Test_Uint16_UintPtr(t *testing.T) {
	var i uint16
	var ii *uint
	ii = Uint16_UintPtr(i)
	if UintPtr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_UintPtr")
	}
}
func Test_UintPtr_Uint16(t *testing.T) {
	var i *uint
	var ii uint16
	ii = UintPtr_Uint16(i, UseDefaultEmpty)
	if Uint16_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint16")
	}
}
func Test_UintPtr_Uint16Ptr(t *testing.T) {
	var i *uint
	var ii *uint16
	ii = UintPtr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Uint(t *testing.T) {
	var i *uint16
	var ii uint
	ii = Uint16Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint")
	}
}
func Test_Uint16Ptr_UintPtr(t *testing.T) {
	var i *uint16
	var ii *uint
	ii = Uint16Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_UintPtr")
	}
}

func Test_UintSlice_Uint16Slice(t *testing.T) {
	var i []uint
	var ii []uint16
	ii = UintSlice_Uint16Slice(i)
	if len(Uint16Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Uint16Slice")
	}
}
func Test_UintSlice_Uint16SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]uint16
	ii = UintSlice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Uint16SlicePtr")
	}
}
func Test_Uint16Slice_UintSlice(t *testing.T) {
	var i []uint16
	var ii []uint
	ii = Uint16Slice_UintSlice(i)
	if len(UintSlice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_UintSlice")
	}
}
func Test_Uint16Slice_UintSlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]uint
	ii = Uint16Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Uint16Slice(t *testing.T) {
	var i *[]uint
	var ii []uint16
	ii = UintSlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint16Slice")
	}
}
func Test_UintSlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]uint16
	ii = UintSlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_UintSlice(t *testing.T) {
	var i *[]uint16
	var ii []uint
	ii = Uint16SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_UintSlice")
	}
}
func Test_Uint16SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]uint
	ii = Uint16SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_UintSlicePtr")
	}
}

func Test_Uint_Uint32(t *testing.T) {
	var i uint
	var ii uint32
	ii = Uint_Uint32(i)
	if Uint32_Uint(ii) != i {
		t.Errorf("error in Uint_Uint32")
	}
}
func Test_Uint_Uint32Ptr(t *testing.T) {
	var i uint
	var ii *uint32
	ii = Uint_Uint32Ptr(i)
	if Uint32Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Uint32Ptr")
	}
}
func Test_Uint32_Uint(t *testing.T) {
	var i uint32
	var ii uint
	ii = Uint32_Uint(i)
	if Uint_Uint32(ii) != i {
		t.Errorf("error in Uint32_Uint")
	}
}
func Test_Uint32_UintPtr(t *testing.T) {
	var i uint32
	var ii *uint
	ii = Uint32_UintPtr(i)
	if UintPtr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_UintPtr")
	}
}
func Test_UintPtr_Uint32(t *testing.T) {
	var i *uint
	var ii uint32
	ii = UintPtr_Uint32(i, UseDefaultEmpty)
	if Uint32_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint32")
	}
}
func Test_UintPtr_Uint32Ptr(t *testing.T) {
	var i *uint
	var ii *uint32
	ii = UintPtr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Uint(t *testing.T) {
	var i *uint32
	var ii uint
	ii = Uint32Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint")
	}
}
func Test_Uint32Ptr_UintPtr(t *testing.T) {
	var i *uint32
	var ii *uint
	ii = Uint32Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_UintPtr")
	}
}

func Test_UintSlice_Uint32Slice(t *testing.T) {
	var i []uint
	var ii []uint32
	ii = UintSlice_Uint32Slice(i)
	if len(Uint32Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Uint32Slice")
	}
}
func Test_UintSlice_Uint32SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]uint32
	ii = UintSlice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Uint32SlicePtr")
	}
}
func Test_Uint32Slice_UintSlice(t *testing.T) {
	var i []uint32
	var ii []uint
	ii = Uint32Slice_UintSlice(i)
	if len(UintSlice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_UintSlice")
	}
}
func Test_Uint32Slice_UintSlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]uint
	ii = Uint32Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Uint32Slice(t *testing.T) {
	var i *[]uint
	var ii []uint32
	ii = UintSlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint32Slice")
	}
}
func Test_UintSlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]uint32
	ii = UintSlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_UintSlice(t *testing.T) {
	var i *[]uint32
	var ii []uint
	ii = Uint32SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_UintSlice")
	}
}
func Test_Uint32SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]uint
	ii = Uint32SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_UintSlicePtr")
	}
}

func Test_Uint_Uint64(t *testing.T) {
	var i uint
	var ii uint64
	ii = Uint_Uint64(i)
	if Uint64_Uint(ii) != i {
		t.Errorf("error in Uint_Uint64")
	}
}
func Test_Uint_Uint64Ptr(t *testing.T) {
	var i uint
	var ii *uint64
	ii = Uint_Uint64Ptr(i)
	if Uint64Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Uint64Ptr")
	}
}
func Test_Uint64_Uint(t *testing.T) {
	var i uint64
	var ii uint
	ii = Uint64_Uint(i)
	if Uint_Uint64(ii) != i {
		t.Errorf("error in Uint64_Uint")
	}
}
func Test_Uint64_UintPtr(t *testing.T) {
	var i uint64
	var ii *uint
	ii = Uint64_UintPtr(i)
	if UintPtr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_UintPtr")
	}
}
func Test_UintPtr_Uint64(t *testing.T) {
	var i *uint
	var ii uint64
	ii = UintPtr_Uint64(i, UseDefaultEmpty)
	if Uint64_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint64")
	}
}
func Test_UintPtr_Uint64Ptr(t *testing.T) {
	var i *uint
	var ii *uint64
	ii = UintPtr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Uint(t *testing.T) {
	var i *uint64
	var ii uint
	ii = Uint64Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint")
	}
}
func Test_Uint64Ptr_UintPtr(t *testing.T) {
	var i *uint64
	var ii *uint
	ii = Uint64Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_UintPtr")
	}
}

func Test_UintSlice_Uint64Slice(t *testing.T) {
	var i []uint
	var ii []uint64
	ii = UintSlice_Uint64Slice(i)
	if len(Uint64Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Uint64Slice")
	}
}
func Test_UintSlice_Uint64SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]uint64
	ii = UintSlice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_UintSlice(t *testing.T) {
	var i []uint64
	var ii []uint
	ii = Uint64Slice_UintSlice(i)
	if len(UintSlice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_UintSlice")
	}
}
func Test_Uint64Slice_UintSlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]uint
	ii = Uint64Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Uint64Slice(t *testing.T) {
	var i *[]uint
	var ii []uint64
	ii = UintSlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint64Slice")
	}
}
func Test_UintSlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]uint64
	ii = UintSlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_UintSlice(t *testing.T) {
	var i *[]uint64
	var ii []uint
	ii = Uint64SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_UintSlice")
	}
}
func Test_Uint64SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]uint
	ii = Uint64SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_UintSlicePtr")
	}
}

func Test_Uint_Uint8(t *testing.T) {
	var i uint
	var ii uint8
	ii = Uint_Uint8(i)
	if Uint8_Uint(ii) != i {
		t.Errorf("error in Uint_Uint8")
	}
}
func Test_Uint_Uint8Ptr(t *testing.T) {
	var i uint
	var ii *uint8
	ii = Uint_Uint8Ptr(i)
	if Uint8Ptr_Uint(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint_Uint8Ptr")
	}
}
func Test_Uint8_Uint(t *testing.T) {
	var i uint8
	var ii uint
	ii = Uint8_Uint(i)
	if Uint_Uint8(ii) != i {
		t.Errorf("error in Uint8_Uint")
	}
}
func Test_Uint8_UintPtr(t *testing.T) {
	var i uint8
	var ii *uint
	ii = Uint8_UintPtr(i)
	if UintPtr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_UintPtr")
	}
}
func Test_UintPtr_Uint8(t *testing.T) {
	var i *uint
	var ii uint8
	ii = UintPtr_Uint8(i, UseDefaultEmpty)
	if Uint8_Uint(ii) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint8")
	}
}
func Test_UintPtr_Uint8Ptr(t *testing.T) {
	var i *uint
	var ii *uint8
	ii = UintPtr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Uint(ii, UseDefaultEmpty) != UintPtr_Uint(i, UseDefaultEmpty) {
		t.Errorf("error in UintPtr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Uint(t *testing.T) {
	var i *uint8
	var ii uint
	ii = Uint8Ptr_Uint(i, UseDefaultEmpty)
	if Uint_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint")
	}
}
func Test_Uint8Ptr_UintPtr(t *testing.T) {
	var i *uint8
	var ii *uint
	ii = Uint8Ptr_UintPtr(i, UseDefaultEmpty)
	if UintPtr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_UintPtr")
	}
}

func Test_UintSlice_Uint8Slice(t *testing.T) {
	var i []uint
	var ii []uint8
	ii = UintSlice_Uint8Slice(i)
	if len(Uint8Slice_UintSlice(ii)) != len(i) {
		t.Errorf("error in UintSlice_Uint8Slice")
	}
}
func Test_UintSlice_Uint8SlicePtr(t *testing.T) {
	var i []uint
	var ii *[]uint8
	ii = UintSlice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in UintSlice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_UintSlice(t *testing.T) {
	var i []uint8
	var ii []uint
	ii = Uint8Slice_UintSlice(i)
	if len(UintSlice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_UintSlice")
	}
}
func Test_Uint8Slice_UintSlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]uint
	ii = Uint8Slice_UintSlicePtr(i)
	if len(UintSlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_UintSlicePtr")
	}
}
func Test_UintSlicePtr_Uint8Slice(t *testing.T) {
	var i *[]uint
	var ii []uint8
	ii = UintSlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_UintSlice(ii)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint8Slice")
	}
}
func Test_UintSlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]uint
	var ii *[]uint8
	ii = UintSlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_UintSlice(ii, UseDefaultEmpty)) != len(UintSlicePtr_UintSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in UintSlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_UintSlice(t *testing.T) {
	var i *[]uint8
	var ii []uint
	ii = Uint8SlicePtr_UintSlice(i, UseDefaultEmpty)
	if len(UintSlice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_UintSlice")
	}
}
func Test_Uint8SlicePtr_UintSlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]uint
	ii = Uint8SlicePtr_UintSlicePtr(i, UseDefaultEmpty)
	if len(UintSlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_UintSlicePtr")
	}
}

func Test_Uint16_Uint32(t *testing.T) {
	var i uint16
	var ii uint32
	ii = Uint16_Uint32(i)
	if Uint32_Uint16(ii) != i {
		t.Errorf("error in Uint16_Uint32")
	}
}
func Test_Uint16_Uint32Ptr(t *testing.T) {
	var i uint16
	var ii *uint32
	ii = Uint16_Uint32Ptr(i)
	if Uint32Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Uint32Ptr")
	}
}
func Test_Uint32_Uint16(t *testing.T) {
	var i uint32
	var ii uint16
	ii = Uint32_Uint16(i)
	if Uint16_Uint32(ii) != i {
		t.Errorf("error in Uint32_Uint16")
	}
}
func Test_Uint32_Uint16Ptr(t *testing.T) {
	var i uint32
	var ii *uint16
	ii = Uint32_Uint16Ptr(i)
	if Uint16Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Uint32(t *testing.T) {
	var i *uint16
	var ii uint32
	ii = Uint16Ptr_Uint32(i, UseDefaultEmpty)
	if Uint32_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint32")
	}
}
func Test_Uint16Ptr_Uint32Ptr(t *testing.T) {
	var i *uint16
	var ii *uint32
	ii = Uint16Ptr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Uint16(t *testing.T) {
	var i *uint32
	var ii uint16
	ii = Uint32Ptr_Uint16(i, UseDefaultEmpty)
	if Uint16_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint16")
	}
}
func Test_Uint32Ptr_Uint16Ptr(t *testing.T) {
	var i *uint32
	var ii *uint16
	ii = Uint32Ptr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint16Ptr")
	}
}

func Test_Uint16Slice_Uint32Slice(t *testing.T) {
	var i []uint16
	var ii []uint32
	ii = Uint16Slice_Uint32Slice(i)
	if len(Uint32Slice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_Uint32Slice")
	}
}
func Test_Uint16Slice_Uint32SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]uint32
	ii = Uint16Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Uint32SlicePtr")
	}
}
func Test_Uint32Slice_Uint16Slice(t *testing.T) {
	var i []uint32
	var ii []uint16
	ii = Uint32Slice_Uint16Slice(i)
	if len(Uint16Slice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_Uint16Slice")
	}
}
func Test_Uint32Slice_Uint16SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]uint16
	ii = Uint32Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]uint16
	var ii []uint32
	ii = Uint16SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Uint32Slice")
	}
}
func Test_Uint16SlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]uint32
	ii = Uint16SlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]uint32
	var ii []uint16
	ii = Uint32SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Uint16Slice")
	}
}
func Test_Uint32SlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]uint16
	ii = Uint32SlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Uint16SlicePtr")
	}
}

func Test_Uint16_Uint64(t *testing.T) {
	var i uint16
	var ii uint64
	ii = Uint16_Uint64(i)
	if Uint64_Uint16(ii) != i {
		t.Errorf("error in Uint16_Uint64")
	}
}
func Test_Uint16_Uint64Ptr(t *testing.T) {
	var i uint16
	var ii *uint64
	ii = Uint16_Uint64Ptr(i)
	if Uint64Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Uint64Ptr")
	}
}
func Test_Uint64_Uint16(t *testing.T) {
	var i uint64
	var ii uint16
	ii = Uint64_Uint16(i)
	if Uint16_Uint64(ii) != i {
		t.Errorf("error in Uint64_Uint16")
	}
}
func Test_Uint64_Uint16Ptr(t *testing.T) {
	var i uint64
	var ii *uint16
	ii = Uint64_Uint16Ptr(i)
	if Uint16Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Uint64(t *testing.T) {
	var i *uint16
	var ii uint64
	ii = Uint16Ptr_Uint64(i, UseDefaultEmpty)
	if Uint64_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint64")
	}
}
func Test_Uint16Ptr_Uint64Ptr(t *testing.T) {
	var i *uint16
	var ii *uint64
	ii = Uint16Ptr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Uint16(t *testing.T) {
	var i *uint64
	var ii uint16
	ii = Uint64Ptr_Uint16(i, UseDefaultEmpty)
	if Uint16_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint16")
	}
}
func Test_Uint64Ptr_Uint16Ptr(t *testing.T) {
	var i *uint64
	var ii *uint16
	ii = Uint64Ptr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint16Ptr")
	}
}

func Test_Uint16Slice_Uint64Slice(t *testing.T) {
	var i []uint16
	var ii []uint64
	ii = Uint16Slice_Uint64Slice(i)
	if len(Uint64Slice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_Uint64Slice")
	}
}
func Test_Uint16Slice_Uint64SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]uint64
	ii = Uint16Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_Uint16Slice(t *testing.T) {
	var i []uint64
	var ii []uint16
	ii = Uint64Slice_Uint16Slice(i)
	if len(Uint16Slice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_Uint16Slice")
	}
}
func Test_Uint64Slice_Uint16SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]uint16
	ii = Uint64Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]uint16
	var ii []uint64
	ii = Uint16SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Uint64Slice")
	}
}
func Test_Uint16SlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]uint64
	ii = Uint16SlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]uint64
	var ii []uint16
	ii = Uint64SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Uint16Slice")
	}
}
func Test_Uint64SlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]uint16
	ii = Uint64SlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Uint16SlicePtr")
	}
}

func Test_Uint16_Uint8(t *testing.T) {
	var i uint16
	var ii uint8
	ii = Uint16_Uint8(i)
	if Uint8_Uint16(ii) != i {
		t.Errorf("error in Uint16_Uint8")
	}
}
func Test_Uint16_Uint8Ptr(t *testing.T) {
	var i uint16
	var ii *uint8
	ii = Uint16_Uint8Ptr(i)
	if Uint8Ptr_Uint16(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint16_Uint8Ptr")
	}
}
func Test_Uint8_Uint16(t *testing.T) {
	var i uint8
	var ii uint16
	ii = Uint8_Uint16(i)
	if Uint16_Uint8(ii) != i {
		t.Errorf("error in Uint8_Uint16")
	}
}
func Test_Uint8_Uint16Ptr(t *testing.T) {
	var i uint8
	var ii *uint16
	ii = Uint8_Uint16Ptr(i)
	if Uint16Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Uint16Ptr")
	}
}
func Test_Uint16Ptr_Uint8(t *testing.T) {
	var i *uint16
	var ii uint8
	ii = Uint16Ptr_Uint8(i, UseDefaultEmpty)
	if Uint8_Uint16(ii) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint8")
	}
}
func Test_Uint16Ptr_Uint8Ptr(t *testing.T) {
	var i *uint16
	var ii *uint8
	ii = Uint16Ptr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Uint16(ii, UseDefaultEmpty) != Uint16Ptr_Uint16(i, UseDefaultEmpty) {
		t.Errorf("error in Uint16Ptr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Uint16(t *testing.T) {
	var i *uint8
	var ii uint16
	ii = Uint8Ptr_Uint16(i, UseDefaultEmpty)
	if Uint16_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint16")
	}
}
func Test_Uint8Ptr_Uint16Ptr(t *testing.T) {
	var i *uint8
	var ii *uint16
	ii = Uint8Ptr_Uint16Ptr(i, UseDefaultEmpty)
	if Uint16Ptr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint16Ptr")
	}
}

func Test_Uint16Slice_Uint8Slice(t *testing.T) {
	var i []uint16
	var ii []uint8
	ii = Uint16Slice_Uint8Slice(i)
	if len(Uint8Slice_Uint16Slice(ii)) != len(i) {
		t.Errorf("error in Uint16Slice_Uint8Slice")
	}
}
func Test_Uint16Slice_Uint8SlicePtr(t *testing.T) {
	var i []uint16
	var ii *[]uint8
	ii = Uint16Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint16Slice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_Uint16Slice(t *testing.T) {
	var i []uint8
	var ii []uint16
	ii = Uint8Slice_Uint16Slice(i)
	if len(Uint16Slice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_Uint16Slice")
	}
}
func Test_Uint8Slice_Uint16SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]uint16
	ii = Uint8Slice_Uint16SlicePtr(i)
	if len(Uint16SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Uint16SlicePtr")
	}
}
func Test_Uint16SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]uint16
	var ii []uint8
	ii = Uint16SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_Uint16Slice(ii)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Uint8Slice")
	}
}
func Test_Uint16SlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]uint16
	var ii *[]uint8
	ii = Uint16SlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_Uint16Slice(ii, UseDefaultEmpty)) != len(Uint16SlicePtr_Uint16Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint16SlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_Uint16Slice(t *testing.T) {
	var i *[]uint8
	var ii []uint16
	ii = Uint8SlicePtr_Uint16Slice(i, UseDefaultEmpty)
	if len(Uint16Slice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Uint16Slice")
	}
}
func Test_Uint8SlicePtr_Uint16SlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]uint16
	ii = Uint8SlicePtr_Uint16SlicePtr(i, UseDefaultEmpty)
	if len(Uint16SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Uint16SlicePtr")
	}
}

func Test_Uint32_Uint64(t *testing.T) {
	var i uint32
	var ii uint64
	ii = Uint32_Uint64(i)
	if Uint64_Uint32(ii) != i {
		t.Errorf("error in Uint32_Uint64")
	}
}
func Test_Uint32_Uint64Ptr(t *testing.T) {
	var i uint32
	var ii *uint64
	ii = Uint32_Uint64Ptr(i)
	if Uint64Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Uint64Ptr")
	}
}
func Test_Uint64_Uint32(t *testing.T) {
	var i uint64
	var ii uint32
	ii = Uint64_Uint32(i)
	if Uint32_Uint64(ii) != i {
		t.Errorf("error in Uint64_Uint32")
	}
}
func Test_Uint64_Uint32Ptr(t *testing.T) {
	var i uint64
	var ii *uint32
	ii = Uint64_Uint32Ptr(i)
	if Uint32Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Uint64(t *testing.T) {
	var i *uint32
	var ii uint64
	ii = Uint32Ptr_Uint64(i, UseDefaultEmpty)
	if Uint64_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint64")
	}
}
func Test_Uint32Ptr_Uint64Ptr(t *testing.T) {
	var i *uint32
	var ii *uint64
	ii = Uint32Ptr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Uint32(t *testing.T) {
	var i *uint64
	var ii uint32
	ii = Uint64Ptr_Uint32(i, UseDefaultEmpty)
	if Uint32_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint32")
	}
}
func Test_Uint64Ptr_Uint32Ptr(t *testing.T) {
	var i *uint64
	var ii *uint32
	ii = Uint64Ptr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint32Ptr")
	}
}

func Test_Uint32Slice_Uint64Slice(t *testing.T) {
	var i []uint32
	var ii []uint64
	ii = Uint32Slice_Uint64Slice(i)
	if len(Uint64Slice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_Uint64Slice")
	}
}
func Test_Uint32Slice_Uint64SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]uint64
	ii = Uint32Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Uint64SlicePtr")
	}
}
func Test_Uint64Slice_Uint32Slice(t *testing.T) {
	var i []uint64
	var ii []uint32
	ii = Uint64Slice_Uint32Slice(i)
	if len(Uint32Slice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_Uint32Slice")
	}
}
func Test_Uint64Slice_Uint32SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]uint32
	ii = Uint64Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]uint32
	var ii []uint64
	ii = Uint32SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Uint64Slice")
	}
}
func Test_Uint32SlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]uint64
	ii = Uint32SlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]uint64
	var ii []uint32
	ii = Uint64SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Uint32Slice")
	}
}
func Test_Uint64SlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]uint32
	ii = Uint64SlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Uint32SlicePtr")
	}
}

func Test_Uint32_Uint8(t *testing.T) {
	var i uint32
	var ii uint8
	ii = Uint32_Uint8(i)
	if Uint8_Uint32(ii) != i {
		t.Errorf("error in Uint32_Uint8")
	}
}
func Test_Uint32_Uint8Ptr(t *testing.T) {
	var i uint32
	var ii *uint8
	ii = Uint32_Uint8Ptr(i)
	if Uint8Ptr_Uint32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint32_Uint8Ptr")
	}
}
func Test_Uint8_Uint32(t *testing.T) {
	var i uint8
	var ii uint32
	ii = Uint8_Uint32(i)
	if Uint32_Uint8(ii) != i {
		t.Errorf("error in Uint8_Uint32")
	}
}
func Test_Uint8_Uint32Ptr(t *testing.T) {
	var i uint8
	var ii *uint32
	ii = Uint8_Uint32Ptr(i)
	if Uint32Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Uint32Ptr")
	}
}
func Test_Uint32Ptr_Uint8(t *testing.T) {
	var i *uint32
	var ii uint8
	ii = Uint32Ptr_Uint8(i, UseDefaultEmpty)
	if Uint8_Uint32(ii) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint8")
	}
}
func Test_Uint32Ptr_Uint8Ptr(t *testing.T) {
	var i *uint32
	var ii *uint8
	ii = Uint32Ptr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Uint32(ii, UseDefaultEmpty) != Uint32Ptr_Uint32(i, UseDefaultEmpty) {
		t.Errorf("error in Uint32Ptr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Uint32(t *testing.T) {
	var i *uint8
	var ii uint32
	ii = Uint8Ptr_Uint32(i, UseDefaultEmpty)
	if Uint32_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint32")
	}
}
func Test_Uint8Ptr_Uint32Ptr(t *testing.T) {
	var i *uint8
	var ii *uint32
	ii = Uint8Ptr_Uint32Ptr(i, UseDefaultEmpty)
	if Uint32Ptr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint32Ptr")
	}
}

func Test_Uint32Slice_Uint8Slice(t *testing.T) {
	var i []uint32
	var ii []uint8
	ii = Uint32Slice_Uint8Slice(i)
	if len(Uint8Slice_Uint32Slice(ii)) != len(i) {
		t.Errorf("error in Uint32Slice_Uint8Slice")
	}
}
func Test_Uint32Slice_Uint8SlicePtr(t *testing.T) {
	var i []uint32
	var ii *[]uint8
	ii = Uint32Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint32Slice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_Uint32Slice(t *testing.T) {
	var i []uint8
	var ii []uint32
	ii = Uint8Slice_Uint32Slice(i)
	if len(Uint32Slice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_Uint32Slice")
	}
}
func Test_Uint8Slice_Uint32SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]uint32
	ii = Uint8Slice_Uint32SlicePtr(i)
	if len(Uint32SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Uint32SlicePtr")
	}
}
func Test_Uint32SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]uint32
	var ii []uint8
	ii = Uint32SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_Uint32Slice(ii)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Uint8Slice")
	}
}
func Test_Uint32SlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]uint32
	var ii *[]uint8
	ii = Uint32SlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_Uint32Slice(ii, UseDefaultEmpty)) != len(Uint32SlicePtr_Uint32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint32SlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_Uint32Slice(t *testing.T) {
	var i *[]uint8
	var ii []uint32
	ii = Uint8SlicePtr_Uint32Slice(i, UseDefaultEmpty)
	if len(Uint32Slice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Uint32Slice")
	}
}
func Test_Uint8SlicePtr_Uint32SlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]uint32
	ii = Uint8SlicePtr_Uint32SlicePtr(i, UseDefaultEmpty)
	if len(Uint32SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Uint32SlicePtr")
	}
}

func Test_Uint64_Uint8(t *testing.T) {
	var i uint64
	var ii uint8
	ii = Uint64_Uint8(i)
	if Uint8_Uint64(ii) != i {
		t.Errorf("error in Uint64_Uint8")
	}
}
func Test_Uint64_Uint8Ptr(t *testing.T) {
	var i uint64
	var ii *uint8
	ii = Uint64_Uint8Ptr(i)
	if Uint8Ptr_Uint64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint64_Uint8Ptr")
	}
}
func Test_Uint8_Uint64(t *testing.T) {
	var i uint8
	var ii uint64
	ii = Uint8_Uint64(i)
	if Uint64_Uint8(ii) != i {
		t.Errorf("error in Uint8_Uint64")
	}
}
func Test_Uint8_Uint64Ptr(t *testing.T) {
	var i uint8
	var ii *uint64
	ii = Uint8_Uint64Ptr(i)
	if Uint64Ptr_Uint8(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Uint8_Uint64Ptr")
	}
}
func Test_Uint64Ptr_Uint8(t *testing.T) {
	var i *uint64
	var ii uint8
	ii = Uint64Ptr_Uint8(i, UseDefaultEmpty)
	if Uint8_Uint64(ii) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint8")
	}
}
func Test_Uint64Ptr_Uint8Ptr(t *testing.T) {
	var i *uint64
	var ii *uint8
	ii = Uint64Ptr_Uint8Ptr(i, UseDefaultEmpty)
	if Uint8Ptr_Uint64(ii, UseDefaultEmpty) != Uint64Ptr_Uint64(i, UseDefaultEmpty) {
		t.Errorf("error in Uint64Ptr_Uint8Ptr")
	}
}
func Test_Uint8Ptr_Uint64(t *testing.T) {
	var i *uint8
	var ii uint64
	ii = Uint8Ptr_Uint64(i, UseDefaultEmpty)
	if Uint64_Uint8(ii) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint64")
	}
}
func Test_Uint8Ptr_Uint64Ptr(t *testing.T) {
	var i *uint8
	var ii *uint64
	ii = Uint8Ptr_Uint64Ptr(i, UseDefaultEmpty)
	if Uint64Ptr_Uint8(ii, UseDefaultEmpty) != Uint8Ptr_Uint8(i, UseDefaultEmpty) {
		t.Errorf("error in Uint8Ptr_Uint64Ptr")
	}
}

func Test_Uint64Slice_Uint8Slice(t *testing.T) {
	var i []uint64
	var ii []uint8
	ii = Uint64Slice_Uint8Slice(i)
	if len(Uint8Slice_Uint64Slice(ii)) != len(i) {
		t.Errorf("error in Uint64Slice_Uint8Slice")
	}
}
func Test_Uint64Slice_Uint8SlicePtr(t *testing.T) {
	var i []uint64
	var ii *[]uint8
	ii = Uint64Slice_Uint8SlicePtr(i)
	if len(Uint8SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint64Slice_Uint8SlicePtr")
	}
}
func Test_Uint8Slice_Uint64Slice(t *testing.T) {
	var i []uint8
	var ii []uint64
	ii = Uint8Slice_Uint64Slice(i)
	if len(Uint64Slice_Uint8Slice(ii)) != len(i) {
		t.Errorf("error in Uint8Slice_Uint64Slice")
	}
}
func Test_Uint8Slice_Uint64SlicePtr(t *testing.T) {
	var i []uint8
	var ii *[]uint64
	ii = Uint8Slice_Uint64SlicePtr(i)
	if len(Uint64SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Uint8Slice_Uint64SlicePtr")
	}
}
func Test_Uint64SlicePtr_Uint8Slice(t *testing.T) {
	var i *[]uint64
	var ii []uint8
	ii = Uint64SlicePtr_Uint8Slice(i, UseDefaultEmpty)
	if len(Uint8Slice_Uint64Slice(ii)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Uint8Slice")
	}
}
func Test_Uint64SlicePtr_Uint8SlicePtr(t *testing.T) {
	var i *[]uint64
	var ii *[]uint8
	ii = Uint64SlicePtr_Uint8SlicePtr(i, UseDefaultEmpty)
	if len(Uint8SlicePtr_Uint64Slice(ii, UseDefaultEmpty)) != len(Uint64SlicePtr_Uint64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint64SlicePtr_Uint8SlicePtr")
	}
}
func Test_Uint8SlicePtr_Uint64Slice(t *testing.T) {
	var i *[]uint8
	var ii []uint64
	ii = Uint8SlicePtr_Uint64Slice(i, UseDefaultEmpty)
	if len(Uint64Slice_Uint8Slice(ii)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Uint64Slice")
	}
}
func Test_Uint8SlicePtr_Uint64SlicePtr(t *testing.T) {
	var i *[]uint8
	var ii *[]uint64
	ii = Uint8SlicePtr_Uint64SlicePtr(i, UseDefaultEmpty)
	if len(Uint64SlicePtr_Uint8Slice(ii, UseDefaultEmpty)) != len(Uint8SlicePtr_Uint8Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Uint8SlicePtr_Uint64SlicePtr")
	}
}

func Test_Float32_Float32Ptr(t *testing.T) {
	var i float32
	var ii *float32
	ii = Float32_Float32Ptr(i)
	if Float32Ptr_Float32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Float32_Float32Ptr")
	}
}
func Test_Float32Ptr_Float32(t *testing.T) {
	var i *float32
	var ii float32
	ii = Float32Ptr_Float32(i, UseDefaultEmpty)
	if (ii) != Float32Ptr_Float32(i, UseDefaultEmpty) {
		t.Errorf("error in Float32Ptr_Float32")
	}
}

func Test_Float32Slice_Float32SlicePtr(t *testing.T) {
	var i []float32
	var ii *[]float32
	ii = Float32Slice_Float32SlicePtr(i)
	if len(Float32SlicePtr_Float32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Float32Slice_Float32SlicePtr")
	}
}
func Test_Float32SlicePtr_Float32Slice(t *testing.T) {
	var i *[]float32
	var ii []float32
	ii = Float32SlicePtr_Float32Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Float32SlicePtr_Float32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Float32SlicePtr_Float32Slice")
	}
}

func Test_Float32_Float64(t *testing.T) {
	var i float32
	var ii float64
	ii = Float32_Float64(i)
	if Float64_Float32(ii) != i {
		t.Errorf("error in Float32_Float64")
	}
}
func Test_Float32_Float64Ptr(t *testing.T) {
	var i float32
	var ii *float64
	ii = Float32_Float64Ptr(i)
	if Float64Ptr_Float32(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Float32_Float64Ptr")
	}
}
func Test_Float64_Float32(t *testing.T) {
	var i float64
	var ii float32
	ii = Float64_Float32(i)
	if Float32_Float64(ii) != i {
		t.Errorf("error in Float64_Float32")
	}
}
func Test_Float64_Float32Ptr(t *testing.T) {
	var i float64
	var ii *float32
	ii = Float64_Float32Ptr(i)
	if Float32Ptr_Float64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Float64_Float32Ptr")
	}
}
func Test_Float64_Float64Ptr(t *testing.T) {
	var i float64
	var ii *float64
	ii = Float64_Float64Ptr(i)
	if Float64Ptr_Float64(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Float64_Float64Ptr")
	}
}
func Test_Float32Ptr_Float64(t *testing.T) {
	var i *float32
	var ii float64
	ii = Float32Ptr_Float64(i, UseDefaultEmpty)
	if Float64_Float32(ii) != Float32Ptr_Float32(i, UseDefaultEmpty) {
		t.Errorf("error in Float32Ptr_Float64")
	}
}
func Test_Float32Ptr_Float64Ptr(t *testing.T) {
	var i *float32
	var ii *float64
	ii = Float32Ptr_Float64Ptr(i, UseDefaultEmpty)
	if Float64Ptr_Float32(ii, UseDefaultEmpty) != Float32Ptr_Float32(i, UseDefaultEmpty) {
		t.Errorf("error in Float32Ptr_Float64Ptr")
	}
}
func Test_Float64Ptr_Float32(t *testing.T) {
	var i *float64
	var ii float32
	ii = Float64Ptr_Float32(i, UseDefaultEmpty)
	if Float32_Float64(ii) != Float64Ptr_Float64(i, UseDefaultEmpty) {
		t.Errorf("error in Float64Ptr_Float32")
	}
}
func Test_Float64Ptr_Float64(t *testing.T) {
	var i *float64
	var ii float64
	ii = Float64Ptr_Float64(i, UseDefaultEmpty)
	if (ii) != Float64Ptr_Float64(i, UseDefaultEmpty) {
		t.Errorf("error in Float64Ptr_Float64")
	}
}
func Test_Float64Ptr_Float32Ptr(t *testing.T) {
	var i *float64
	var ii *float32
	ii = Float64Ptr_Float32Ptr(i, UseDefaultEmpty)
	if Float32Ptr_Float64(ii, UseDefaultEmpty) != Float64Ptr_Float64(i, UseDefaultEmpty) {
		t.Errorf("error in Float64Ptr_Float32Ptr")
	}
}

func Test_Float32Slice_Float64Slice(t *testing.T) {
	var i []float32
	var ii []float64
	ii = Float32Slice_Float64Slice(i)
	if len(Float64Slice_Float32Slice(ii)) != len(i) {
		t.Errorf("error in Float32Slice_Float64Slice")
	}
}
func Test_Float32Slice_Float64SlicePtr(t *testing.T) {
	var i []float32
	var ii *[]float64
	ii = Float32Slice_Float64SlicePtr(i)
	if len(Float64SlicePtr_Float32Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Float32Slice_Float64SlicePtr")
	}
}
func Test_Float64Slice_Float32Slice(t *testing.T) {
	var i []float64
	var ii []float32
	ii = Float64Slice_Float32Slice(i)
	if len(Float32Slice_Float64Slice(ii)) != len(i) {
		t.Errorf("error in Float64Slice_Float32Slice")
	}
}
func Test_Float64Slice_Float32SlicePtr(t *testing.T) {
	var i []float64
	var ii *[]float32
	ii = Float64Slice_Float32SlicePtr(i)
	if len(Float32SlicePtr_Float64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Float64Slice_Float32SlicePtr")
	}
}
func Test_Float64Slice_Float64SlicePtr(t *testing.T) {
	var i []float64
	var ii *[]float64
	ii = Float64Slice_Float64SlicePtr(i)
	if len(Float64SlicePtr_Float64Slice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in Float64Slice_Float64SlicePtr")
	}
}
func Test_Float32SlicePtr_Float64Slice(t *testing.T) {
	var i *[]float32
	var ii []float64
	ii = Float32SlicePtr_Float64Slice(i, UseDefaultEmpty)
	if len(Float64Slice_Float32Slice(ii)) != len(Float32SlicePtr_Float32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Float32SlicePtr_Float64Slice")
	}
}
func Test_Float32SlicePtr_Float64SlicePtr(t *testing.T) {
	var i *[]float32
	var ii *[]float64
	ii = Float32SlicePtr_Float64SlicePtr(i, UseDefaultEmpty)
	if len(Float64SlicePtr_Float32Slice(ii, UseDefaultEmpty)) != len(Float32SlicePtr_Float32Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Float32SlicePtr_Float64SlicePtr")
	}
}
func Test_Float64SlicePtr_Float32Slice(t *testing.T) {
	var i *[]float64
	var ii []float32
	ii = Float64SlicePtr_Float32Slice(i, UseDefaultEmpty)
	if len(Float32Slice_Float64Slice(ii)) != len(Float64SlicePtr_Float64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Float64SlicePtr_Float32Slice")
	}
}
func Test_Float64SlicePtr_Float64Slice(t *testing.T) {
	var i *[]float64
	var ii []float64
	ii = Float64SlicePtr_Float64Slice(i, UseDefaultEmpty)
	if len((ii)) != len(Float64SlicePtr_Float64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Float64SlicePtr_Float64Slice")
	}
}
func Test_Float64SlicePtr_Float32SlicePtr(t *testing.T) {
	var i *[]float64
	var ii *[]float32
	ii = Float64SlicePtr_Float32SlicePtr(i, UseDefaultEmpty)
	if len(Float32SlicePtr_Float64Slice(ii, UseDefaultEmpty)) != len(Float64SlicePtr_Float64Slice(i, UseDefaultEmpty)) {
		t.Errorf("error in Float64SlicePtr_Float32SlicePtr")
	}
}

func Test_String_StringPtr(t *testing.T) {
	var i string
	var ii *string
	ii = String_StringPtr(i)
	if StringPtr_String(ii, UseDefaultEmpty) != i {
		t.Errorf("error in String_StringPtr")
	}
}
func Test_StringPtr_String(t *testing.T) {
	var i *string
	var ii string
	ii = StringPtr_String(i, UseDefaultEmpty)
	if (ii) != StringPtr_String(i, UseDefaultEmpty) {
		t.Errorf("error in StringPtr_String")
	}
}

func Test_StringSlice_StringSlicePtr(t *testing.T) {
	var i []string
	var ii *[]string
	ii = StringSlice_StringSlicePtr(i)
	if len(StringSlicePtr_StringSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in StringSlice_StringSlicePtr")
	}
}
func Test_StringSlicePtr_StringSlice(t *testing.T) {
	var i *[]string
	var ii []string
	ii = StringSlicePtr_StringSlice(i, UseDefaultEmpty)
	if len((ii)) != len(StringSlicePtr_StringSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in StringSlicePtr_StringSlice")
	}
}

func Test_Bool_BoolPtr(t *testing.T) {
	var i bool
	var ii *bool
	ii = Bool_BoolPtr(i)
	if BoolPtr_Bool(ii, UseDefaultEmpty) != i {
		t.Errorf("error in Bool_BoolPtr")
	}
}
func Test_BoolPtr_Bool(t *testing.T) {
	var i *bool
	var ii bool
	ii = BoolPtr_Bool(i, UseDefaultEmpty)
	if (ii) != BoolPtr_Bool(i, UseDefaultEmpty) {
		t.Errorf("error in BoolPtr_Bool")
	}
}

func Test_BoolSlice_BoolSlicePtr(t *testing.T) {
	var i []bool
	var ii *[]bool
	ii = BoolSlice_BoolSlicePtr(i)
	if len(BoolSlicePtr_BoolSlice(ii, UseDefaultEmpty)) != len(i) {
		t.Errorf("error in BoolSlice_BoolSlicePtr")
	}
}
func Test_BoolSlicePtr_BoolSlice(t *testing.T) {
	var i *[]bool
	var ii []bool
	ii = BoolSlicePtr_BoolSlice(i, UseDefaultEmpty)
	if len((ii)) != len(BoolSlicePtr_BoolSlice(i, UseDefaultEmpty)) {
		t.Errorf("error in BoolSlicePtr_BoolSlice")
	}
}
