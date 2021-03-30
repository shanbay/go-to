package to

func Int_IntPtr(i int) *int { return &i }
func IntPtr_Int(i *int, opt ...Option) int {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(int)
	}
	return *i
}

func IntSlice_IntSlicePtr(i []int) *[]int { return &i }
func IntSlicePtr_IntSlice(i *[]int, opt ...Option) []int {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]int)
	}
	return *i
}

func Int_Int16(i int) int16 {
	return int16(i)
}
func Int_Int16Ptr(i int) *int16 {
	return Int16_Int16Ptr(Int_Int16(i))
}
func Int16_Int(i int16) int {
	return int(i)
}
func Int16_IntPtr(i int16) *int {
	return Int_IntPtr(Int16_Int(i))
}
func Int16_Int16Ptr(i int16) *int16 { return &i }
func IntPtr_Int16(i *int, opt ...Option) int16 {
	return Int_Int16(IntPtr_Int(i, opt...))
}
func IntPtr_Int16Ptr(i *int, opt ...Option) *int16 {
	return Int16_Int16Ptr(Int_Int16(IntPtr_Int(i, opt...)))
}
func Int16Ptr_Int(i *int16, opt ...Option) int {
	return Int16_Int(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Int16(i *int16, opt ...Option) int16 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(int16)
	}
	return *i
}
func Int16Ptr_IntPtr(i *int16, opt ...Option) *int {
	return Int_IntPtr(Int16_Int(Int16Ptr_Int16(i, opt...)))
}

func IntSlice_Int16Slice(i []int) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func IntSlice_Int16SlicePtr(i []int) *[]int16 {
	return Int16Slice_Int16SlicePtr(IntSlice_Int16Slice(i))
}
func Int16Slice_IntSlice(i []int16) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Int16Slice_IntSlicePtr(i []int16) *[]int {
	return IntSlice_IntSlicePtr(Int16Slice_IntSlice(i))
}
func Int16Slice_Int16SlicePtr(i []int16) *[]int16 { return &i }
func IntSlicePtr_Int16Slice(i *[]int, opt ...Option) []int16 {
	return IntSlice_Int16Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Int16SlicePtr(i *[]int, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(IntSlice_Int16Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Int16SlicePtr_IntSlice(i *[]int16, opt ...Option) []int {
	return Int16Slice_IntSlice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Int16Slice(i *[]int16, opt ...Option) []int16 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]int16)
	}
	return *i
}
func Int16SlicePtr_IntSlicePtr(i *[]int16, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Int16Slice_IntSlice(Int16SlicePtr_Int16Slice(i, opt...)))
}

func Int_Int32(i int) int32 {
	return int32(i)
}
func Int_Int32Ptr(i int) *int32 {
	return Int32_Int32Ptr(Int_Int32(i))
}
func Int32_Int(i int32) int {
	return int(i)
}
func Int32_IntPtr(i int32) *int {
	return Int_IntPtr(Int32_Int(i))
}
func Int32_Int32Ptr(i int32) *int32 { return &i }
func IntPtr_Int32(i *int, opt ...Option) int32 {
	return Int_Int32(IntPtr_Int(i, opt...))
}
func IntPtr_Int32Ptr(i *int, opt ...Option) *int32 {
	return Int32_Int32Ptr(Int_Int32(IntPtr_Int(i, opt...)))
}
func Int32Ptr_Int(i *int32, opt ...Option) int {
	return Int32_Int(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Int32(i *int32, opt ...Option) int32 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(int32)
	}
	return *i
}
func Int32Ptr_IntPtr(i *int32, opt ...Option) *int {
	return Int_IntPtr(Int32_Int(Int32Ptr_Int32(i, opt...)))
}

func IntSlice_Int32Slice(i []int) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func IntSlice_Int32SlicePtr(i []int) *[]int32 {
	return Int32Slice_Int32SlicePtr(IntSlice_Int32Slice(i))
}
func Int32Slice_IntSlice(i []int32) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Int32Slice_IntSlicePtr(i []int32) *[]int {
	return IntSlice_IntSlicePtr(Int32Slice_IntSlice(i))
}
func Int32Slice_Int32SlicePtr(i []int32) *[]int32 { return &i }
func IntSlicePtr_Int32Slice(i *[]int, opt ...Option) []int32 {
	return IntSlice_Int32Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Int32SlicePtr(i *[]int, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(IntSlice_Int32Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Int32SlicePtr_IntSlice(i *[]int32, opt ...Option) []int {
	return Int32Slice_IntSlice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Int32Slice(i *[]int32, opt ...Option) []int32 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]int32)
	}
	return *i
}
func Int32SlicePtr_IntSlicePtr(i *[]int32, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Int32Slice_IntSlice(Int32SlicePtr_Int32Slice(i, opt...)))
}

func Int_Int64(i int) int64 {
	return int64(i)
}
func Int_Int64Ptr(i int) *int64 {
	return Int64_Int64Ptr(Int_Int64(i))
}
func Int64_Int(i int64) int {
	return int(i)
}
func Int64_IntPtr(i int64) *int {
	return Int_IntPtr(Int64_Int(i))
}
func Int64_Int64Ptr(i int64) *int64 { return &i }
func IntPtr_Int64(i *int, opt ...Option) int64 {
	return Int_Int64(IntPtr_Int(i, opt...))
}
func IntPtr_Int64Ptr(i *int, opt ...Option) *int64 {
	return Int64_Int64Ptr(Int_Int64(IntPtr_Int(i, opt...)))
}
func Int64Ptr_Int(i *int64, opt ...Option) int {
	return Int64_Int(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Int64(i *int64, opt ...Option) int64 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(int64)
	}
	return *i
}
func Int64Ptr_IntPtr(i *int64, opt ...Option) *int {
	return Int_IntPtr(Int64_Int(Int64Ptr_Int64(i, opt...)))
}

func IntSlice_Int64Slice(i []int) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func IntSlice_Int64SlicePtr(i []int) *[]int64 {
	return Int64Slice_Int64SlicePtr(IntSlice_Int64Slice(i))
}
func Int64Slice_IntSlice(i []int64) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Int64Slice_IntSlicePtr(i []int64) *[]int {
	return IntSlice_IntSlicePtr(Int64Slice_IntSlice(i))
}
func Int64Slice_Int64SlicePtr(i []int64) *[]int64 { return &i }
func IntSlicePtr_Int64Slice(i *[]int, opt ...Option) []int64 {
	return IntSlice_Int64Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Int64SlicePtr(i *[]int, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(IntSlice_Int64Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Int64SlicePtr_IntSlice(i *[]int64, opt ...Option) []int {
	return Int64Slice_IntSlice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Int64Slice(i *[]int64, opt ...Option) []int64 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]int64)
	}
	return *i
}
func Int64SlicePtr_IntSlicePtr(i *[]int64, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Int64Slice_IntSlice(Int64SlicePtr_Int64Slice(i, opt...)))
}

func Int_Int8(i int) int8 {
	return int8(i)
}
func Int_Int8Ptr(i int) *int8 {
	return Int8_Int8Ptr(Int_Int8(i))
}
func Int8_Int(i int8) int {
	return int(i)
}
func Int8_IntPtr(i int8) *int {
	return Int_IntPtr(Int8_Int(i))
}
func Int8_Int8Ptr(i int8) *int8 { return &i }
func IntPtr_Int8(i *int, opt ...Option) int8 {
	return Int_Int8(IntPtr_Int(i, opt...))
}
func IntPtr_Int8Ptr(i *int, opt ...Option) *int8 {
	return Int8_Int8Ptr(Int_Int8(IntPtr_Int(i, opt...)))
}
func Int8Ptr_Int(i *int8, opt ...Option) int {
	return Int8_Int(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Int8(i *int8, opt ...Option) int8 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(int8)
	}
	return *i
}
func Int8Ptr_IntPtr(i *int8, opt ...Option) *int {
	return Int_IntPtr(Int8_Int(Int8Ptr_Int8(i, opt...)))
}

func IntSlice_Int8Slice(i []int) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func IntSlice_Int8SlicePtr(i []int) *[]int8 {
	return Int8Slice_Int8SlicePtr(IntSlice_Int8Slice(i))
}
func Int8Slice_IntSlice(i []int8) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Int8Slice_IntSlicePtr(i []int8) *[]int {
	return IntSlice_IntSlicePtr(Int8Slice_IntSlice(i))
}
func Int8Slice_Int8SlicePtr(i []int8) *[]int8 { return &i }
func IntSlicePtr_Int8Slice(i *[]int, opt ...Option) []int8 {
	return IntSlice_Int8Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Int8SlicePtr(i *[]int, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(IntSlice_Int8Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Int8SlicePtr_IntSlice(i *[]int8, opt ...Option) []int {
	return Int8Slice_IntSlice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Int8Slice(i *[]int8, opt ...Option) []int8 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]int8)
	}
	return *i
}
func Int8SlicePtr_IntSlicePtr(i *[]int8, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Int8Slice_IntSlice(Int8SlicePtr_Int8Slice(i, opt...)))
}

func Int_Uint(i int) uint {
	return uint(i)
}
func Int_UintPtr(i int) *uint {
	return Uint_UintPtr(Int_Uint(i))
}
func Uint_Int(i uint) int {
	return int(i)
}
func Uint_IntPtr(i uint) *int {
	return Int_IntPtr(Uint_Int(i))
}
func Uint_UintPtr(i uint) *uint { return &i }
func IntPtr_Uint(i *int, opt ...Option) uint {
	return Int_Uint(IntPtr_Int(i, opt...))
}
func IntPtr_UintPtr(i *int, opt ...Option) *uint {
	return Uint_UintPtr(Int_Uint(IntPtr_Int(i, opt...)))
}
func UintPtr_Int(i *uint, opt ...Option) int {
	return Uint_Int(UintPtr_Uint(i, opt...))
}
func UintPtr_Uint(i *uint, opt ...Option) uint {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(uint)
	}
	return *i
}
func UintPtr_IntPtr(i *uint, opt ...Option) *int {
	return Int_IntPtr(Uint_Int(UintPtr_Uint(i, opt...)))
}

func IntSlice_UintSlice(i []int) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func IntSlice_UintSlicePtr(i []int) *[]uint {
	return UintSlice_UintSlicePtr(IntSlice_UintSlice(i))
}
func UintSlice_IntSlice(i []uint) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func UintSlice_IntSlicePtr(i []uint) *[]int {
	return IntSlice_IntSlicePtr(UintSlice_IntSlice(i))
}
func UintSlice_UintSlicePtr(i []uint) *[]uint { return &i }
func IntSlicePtr_UintSlice(i *[]int, opt ...Option) []uint {
	return IntSlice_UintSlice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_UintSlicePtr(i *[]int, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(IntSlice_UintSlice(IntSlicePtr_IntSlice(i, opt...)))
}
func UintSlicePtr_IntSlice(i *[]uint, opt ...Option) []int {
	return UintSlice_IntSlice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_UintSlice(i *[]uint, opt ...Option) []uint {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]uint)
	}
	return *i
}
func UintSlicePtr_IntSlicePtr(i *[]uint, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(UintSlice_IntSlice(UintSlicePtr_UintSlice(i, opt...)))
}

func Int_Uint16(i int) uint16 {
	return uint16(i)
}
func Int_Uint16Ptr(i int) *uint16 {
	return Uint16_Uint16Ptr(Int_Uint16(i))
}
func Uint16_Int(i uint16) int {
	return int(i)
}
func Uint16_IntPtr(i uint16) *int {
	return Int_IntPtr(Uint16_Int(i))
}
func Uint16_Uint16Ptr(i uint16) *uint16 { return &i }
func IntPtr_Uint16(i *int, opt ...Option) uint16 {
	return Int_Uint16(IntPtr_Int(i, opt...))
}
func IntPtr_Uint16Ptr(i *int, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Int_Uint16(IntPtr_Int(i, opt...)))
}
func Uint16Ptr_Int(i *uint16, opt ...Option) int {
	return Uint16_Int(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Uint16(i *uint16, opt ...Option) uint16 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(uint16)
	}
	return *i
}
func Uint16Ptr_IntPtr(i *uint16, opt ...Option) *int {
	return Int_IntPtr(Uint16_Int(Uint16Ptr_Uint16(i, opt...)))
}

func IntSlice_Uint16Slice(i []int) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func IntSlice_Uint16SlicePtr(i []int) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(IntSlice_Uint16Slice(i))
}
func Uint16Slice_IntSlice(i []uint16) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Uint16Slice_IntSlicePtr(i []uint16) *[]int {
	return IntSlice_IntSlicePtr(Uint16Slice_IntSlice(i))
}
func Uint16Slice_Uint16SlicePtr(i []uint16) *[]uint16 { return &i }
func IntSlicePtr_Uint16Slice(i *[]int, opt ...Option) []uint16 {
	return IntSlice_Uint16Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Uint16SlicePtr(i *[]int, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(IntSlice_Uint16Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Uint16SlicePtr_IntSlice(i *[]uint16, opt ...Option) []int {
	return Uint16Slice_IntSlice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Uint16Slice(i *[]uint16, opt ...Option) []uint16 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]uint16)
	}
	return *i
}
func Uint16SlicePtr_IntSlicePtr(i *[]uint16, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Uint16Slice_IntSlice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}

func Int_Uint32(i int) uint32 {
	return uint32(i)
}
func Int_Uint32Ptr(i int) *uint32 {
	return Uint32_Uint32Ptr(Int_Uint32(i))
}
func Uint32_Int(i uint32) int {
	return int(i)
}
func Uint32_IntPtr(i uint32) *int {
	return Int_IntPtr(Uint32_Int(i))
}
func Uint32_Uint32Ptr(i uint32) *uint32 { return &i }
func IntPtr_Uint32(i *int, opt ...Option) uint32 {
	return Int_Uint32(IntPtr_Int(i, opt...))
}
func IntPtr_Uint32Ptr(i *int, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Int_Uint32(IntPtr_Int(i, opt...)))
}
func Uint32Ptr_Int(i *uint32, opt ...Option) int {
	return Uint32_Int(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Uint32(i *uint32, opt ...Option) uint32 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(uint32)
	}
	return *i
}
func Uint32Ptr_IntPtr(i *uint32, opt ...Option) *int {
	return Int_IntPtr(Uint32_Int(Uint32Ptr_Uint32(i, opt...)))
}

func IntSlice_Uint32Slice(i []int) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func IntSlice_Uint32SlicePtr(i []int) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(IntSlice_Uint32Slice(i))
}
func Uint32Slice_IntSlice(i []uint32) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Uint32Slice_IntSlicePtr(i []uint32) *[]int {
	return IntSlice_IntSlicePtr(Uint32Slice_IntSlice(i))
}
func Uint32Slice_Uint32SlicePtr(i []uint32) *[]uint32 { return &i }
func IntSlicePtr_Uint32Slice(i *[]int, opt ...Option) []uint32 {
	return IntSlice_Uint32Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Uint32SlicePtr(i *[]int, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(IntSlice_Uint32Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Uint32SlicePtr_IntSlice(i *[]uint32, opt ...Option) []int {
	return Uint32Slice_IntSlice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Uint32Slice(i *[]uint32, opt ...Option) []uint32 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]uint32)
	}
	return *i
}
func Uint32SlicePtr_IntSlicePtr(i *[]uint32, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Uint32Slice_IntSlice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}

func Int_Uint64(i int) uint64 {
	return uint64(i)
}
func Int_Uint64Ptr(i int) *uint64 {
	return Uint64_Uint64Ptr(Int_Uint64(i))
}
func Uint64_Int(i uint64) int {
	return int(i)
}
func Uint64_IntPtr(i uint64) *int {
	return Int_IntPtr(Uint64_Int(i))
}
func Uint64_Uint64Ptr(i uint64) *uint64 { return &i }
func IntPtr_Uint64(i *int, opt ...Option) uint64 {
	return Int_Uint64(IntPtr_Int(i, opt...))
}
func IntPtr_Uint64Ptr(i *int, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Int_Uint64(IntPtr_Int(i, opt...)))
}
func Uint64Ptr_Int(i *uint64, opt ...Option) int {
	return Uint64_Int(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Uint64(i *uint64, opt ...Option) uint64 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(uint64)
	}
	return *i
}
func Uint64Ptr_IntPtr(i *uint64, opt ...Option) *int {
	return Int_IntPtr(Uint64_Int(Uint64Ptr_Uint64(i, opt...)))
}

func IntSlice_Uint64Slice(i []int) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func IntSlice_Uint64SlicePtr(i []int) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(IntSlice_Uint64Slice(i))
}
func Uint64Slice_IntSlice(i []uint64) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Uint64Slice_IntSlicePtr(i []uint64) *[]int {
	return IntSlice_IntSlicePtr(Uint64Slice_IntSlice(i))
}
func Uint64Slice_Uint64SlicePtr(i []uint64) *[]uint64 { return &i }
func IntSlicePtr_Uint64Slice(i *[]int, opt ...Option) []uint64 {
	return IntSlice_Uint64Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Uint64SlicePtr(i *[]int, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(IntSlice_Uint64Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Uint64SlicePtr_IntSlice(i *[]uint64, opt ...Option) []int {
	return Uint64Slice_IntSlice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Uint64Slice(i *[]uint64, opt ...Option) []uint64 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]uint64)
	}
	return *i
}
func Uint64SlicePtr_IntSlicePtr(i *[]uint64, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Uint64Slice_IntSlice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Int_Uint8(i int) uint8 {
	return uint8(i)
}
func Int_Uint8Ptr(i int) *uint8 {
	return Uint8_Uint8Ptr(Int_Uint8(i))
}
func Uint8_Int(i uint8) int {
	return int(i)
}
func Uint8_IntPtr(i uint8) *int {
	return Int_IntPtr(Uint8_Int(i))
}
func Uint8_Uint8Ptr(i uint8) *uint8 { return &i }
func IntPtr_Uint8(i *int, opt ...Option) uint8 {
	return Int_Uint8(IntPtr_Int(i, opt...))
}
func IntPtr_Uint8Ptr(i *int, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Int_Uint8(IntPtr_Int(i, opt...)))
}
func Uint8Ptr_Int(i *uint8, opt ...Option) int {
	return Uint8_Int(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Uint8(i *uint8, opt ...Option) uint8 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(uint8)
	}
	return *i
}
func Uint8Ptr_IntPtr(i *uint8, opt ...Option) *int {
	return Int_IntPtr(Uint8_Int(Uint8Ptr_Uint8(i, opt...)))
}

func IntSlice_Uint8Slice(i []int) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func IntSlice_Uint8SlicePtr(i []int) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(IntSlice_Uint8Slice(i))
}
func Uint8Slice_IntSlice(i []uint8) []int {
	res := []int{}
	for _, item := range i {
		res = append(res, int(item))
	}
	return res
}
func Uint8Slice_IntSlicePtr(i []uint8) *[]int {
	return IntSlice_IntSlicePtr(Uint8Slice_IntSlice(i))
}
func Uint8Slice_Uint8SlicePtr(i []uint8) *[]uint8 { return &i }
func IntSlicePtr_Uint8Slice(i *[]int, opt ...Option) []uint8 {
	return IntSlice_Uint8Slice(IntSlicePtr_IntSlice(i, opt...))
}
func IntSlicePtr_Uint8SlicePtr(i *[]int, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(IntSlice_Uint8Slice(IntSlicePtr_IntSlice(i, opt...)))
}
func Uint8SlicePtr_IntSlice(i *[]uint8, opt ...Option) []int {
	return Uint8Slice_IntSlice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Uint8Slice(i *[]uint8, opt ...Option) []uint8 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]uint8)
	}
	return *i
}
func Uint8SlicePtr_IntSlicePtr(i *[]uint8, opt ...Option) *[]int {
	return IntSlice_IntSlicePtr(Uint8Slice_IntSlice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Int16_Int32(i int16) int32 {
	return int32(i)
}
func Int16_Int32Ptr(i int16) *int32 {
	return Int32_Int32Ptr(Int16_Int32(i))
}
func Int32_Int16(i int32) int16 {
	return int16(i)
}
func Int32_Int16Ptr(i int32) *int16 {
	return Int16_Int16Ptr(Int32_Int16(i))
}
func Int16Ptr_Int32(i *int16, opt ...Option) int32 {
	return Int16_Int32(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Int32Ptr(i *int16, opt ...Option) *int32 {
	return Int32_Int32Ptr(Int16_Int32(Int16Ptr_Int16(i, opt...)))
}
func Int32Ptr_Int16(i *int32, opt ...Option) int16 {
	return Int32_Int16(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Int16Ptr(i *int32, opt ...Option) *int16 {
	return Int16_Int16Ptr(Int32_Int16(Int32Ptr_Int32(i, opt...)))
}

func Int16Slice_Int32Slice(i []int16) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func Int16Slice_Int32SlicePtr(i []int16) *[]int32 {
	return Int32Slice_Int32SlicePtr(Int16Slice_Int32Slice(i))
}
func Int32Slice_Int16Slice(i []int32) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func Int32Slice_Int16SlicePtr(i []int32) *[]int16 {
	return Int16Slice_Int16SlicePtr(Int32Slice_Int16Slice(i))
}
func Int16SlicePtr_Int32Slice(i *[]int16, opt ...Option) []int32 {
	return Int16Slice_Int32Slice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Int32SlicePtr(i *[]int16, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(Int16Slice_Int32Slice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func Int32SlicePtr_Int16Slice(i *[]int32, opt ...Option) []int16 {
	return Int32Slice_Int16Slice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Int16SlicePtr(i *[]int32, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(Int32Slice_Int16Slice(Int32SlicePtr_Int32Slice(i, opt...)))
}

func Int16_Int64(i int16) int64 {
	return int64(i)
}
func Int16_Int64Ptr(i int16) *int64 {
	return Int64_Int64Ptr(Int16_Int64(i))
}
func Int64_Int16(i int64) int16 {
	return int16(i)
}
func Int64_Int16Ptr(i int64) *int16 {
	return Int16_Int16Ptr(Int64_Int16(i))
}
func Int16Ptr_Int64(i *int16, opt ...Option) int64 {
	return Int16_Int64(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Int64Ptr(i *int16, opt ...Option) *int64 {
	return Int64_Int64Ptr(Int16_Int64(Int16Ptr_Int16(i, opt...)))
}
func Int64Ptr_Int16(i *int64, opt ...Option) int16 {
	return Int64_Int16(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Int16Ptr(i *int64, opt ...Option) *int16 {
	return Int16_Int16Ptr(Int64_Int16(Int64Ptr_Int64(i, opt...)))
}

func Int16Slice_Int64Slice(i []int16) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func Int16Slice_Int64SlicePtr(i []int16) *[]int64 {
	return Int64Slice_Int64SlicePtr(Int16Slice_Int64Slice(i))
}
func Int64Slice_Int16Slice(i []int64) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func Int64Slice_Int16SlicePtr(i []int64) *[]int16 {
	return Int16Slice_Int16SlicePtr(Int64Slice_Int16Slice(i))
}
func Int16SlicePtr_Int64Slice(i *[]int16, opt ...Option) []int64 {
	return Int16Slice_Int64Slice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Int64SlicePtr(i *[]int16, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(Int16Slice_Int64Slice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func Int64SlicePtr_Int16Slice(i *[]int64, opt ...Option) []int16 {
	return Int64Slice_Int16Slice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Int16SlicePtr(i *[]int64, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(Int64Slice_Int16Slice(Int64SlicePtr_Int64Slice(i, opt...)))
}

func Int16_Int8(i int16) int8 {
	return int8(i)
}
func Int16_Int8Ptr(i int16) *int8 {
	return Int8_Int8Ptr(Int16_Int8(i))
}
func Int8_Int16(i int8) int16 {
	return int16(i)
}
func Int8_Int16Ptr(i int8) *int16 {
	return Int16_Int16Ptr(Int8_Int16(i))
}
func Int16Ptr_Int8(i *int16, opt ...Option) int8 {
	return Int16_Int8(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Int8Ptr(i *int16, opt ...Option) *int8 {
	return Int8_Int8Ptr(Int16_Int8(Int16Ptr_Int16(i, opt...)))
}
func Int8Ptr_Int16(i *int8, opt ...Option) int16 {
	return Int8_Int16(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Int16Ptr(i *int8, opt ...Option) *int16 {
	return Int16_Int16Ptr(Int8_Int16(Int8Ptr_Int8(i, opt...)))
}

func Int16Slice_Int8Slice(i []int16) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func Int16Slice_Int8SlicePtr(i []int16) *[]int8 {
	return Int8Slice_Int8SlicePtr(Int16Slice_Int8Slice(i))
}
func Int8Slice_Int16Slice(i []int8) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func Int8Slice_Int16SlicePtr(i []int8) *[]int16 {
	return Int16Slice_Int16SlicePtr(Int8Slice_Int16Slice(i))
}
func Int16SlicePtr_Int8Slice(i *[]int16, opt ...Option) []int8 {
	return Int16Slice_Int8Slice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Int8SlicePtr(i *[]int16, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(Int16Slice_Int8Slice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func Int8SlicePtr_Int16Slice(i *[]int8, opt ...Option) []int16 {
	return Int8Slice_Int16Slice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Int16SlicePtr(i *[]int8, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(Int8Slice_Int16Slice(Int8SlicePtr_Int8Slice(i, opt...)))
}

func Int16_Uint(i int16) uint {
	return uint(i)
}
func Int16_UintPtr(i int16) *uint {
	return Uint_UintPtr(Int16_Uint(i))
}
func Uint_Int16(i uint) int16 {
	return int16(i)
}
func Uint_Int16Ptr(i uint) *int16 {
	return Int16_Int16Ptr(Uint_Int16(i))
}
func Int16Ptr_Uint(i *int16, opt ...Option) uint {
	return Int16_Uint(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_UintPtr(i *int16, opt ...Option) *uint {
	return Uint_UintPtr(Int16_Uint(Int16Ptr_Int16(i, opt...)))
}
func UintPtr_Int16(i *uint, opt ...Option) int16 {
	return Uint_Int16(UintPtr_Uint(i, opt...))
}
func UintPtr_Int16Ptr(i *uint, opt ...Option) *int16 {
	return Int16_Int16Ptr(Uint_Int16(UintPtr_Uint(i, opt...)))
}

func Int16Slice_UintSlice(i []int16) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Int16Slice_UintSlicePtr(i []int16) *[]uint {
	return UintSlice_UintSlicePtr(Int16Slice_UintSlice(i))
}
func UintSlice_Int16Slice(i []uint) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func UintSlice_Int16SlicePtr(i []uint) *[]int16 {
	return Int16Slice_Int16SlicePtr(UintSlice_Int16Slice(i))
}
func Int16SlicePtr_UintSlice(i *[]int16, opt ...Option) []uint {
	return Int16Slice_UintSlice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_UintSlicePtr(i *[]int16, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Int16Slice_UintSlice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func UintSlicePtr_Int16Slice(i *[]uint, opt ...Option) []int16 {
	return UintSlice_Int16Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Int16SlicePtr(i *[]uint, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(UintSlice_Int16Slice(UintSlicePtr_UintSlice(i, opt...)))
}

func Int16_Uint16(i int16) uint16 {
	return uint16(i)
}
func Int16_Uint16Ptr(i int16) *uint16 {
	return Uint16_Uint16Ptr(Int16_Uint16(i))
}
func Uint16_Int16(i uint16) int16 {
	return int16(i)
}
func Uint16_Int16Ptr(i uint16) *int16 {
	return Int16_Int16Ptr(Uint16_Int16(i))
}
func Int16Ptr_Uint16(i *int16, opt ...Option) uint16 {
	return Int16_Uint16(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Uint16Ptr(i *int16, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Int16_Uint16(Int16Ptr_Int16(i, opt...)))
}
func Uint16Ptr_Int16(i *uint16, opt ...Option) int16 {
	return Uint16_Int16(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Int16Ptr(i *uint16, opt ...Option) *int16 {
	return Int16_Int16Ptr(Uint16_Int16(Uint16Ptr_Uint16(i, opt...)))
}

func Int16Slice_Uint16Slice(i []int16) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func Int16Slice_Uint16SlicePtr(i []int16) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int16Slice_Uint16Slice(i))
}
func Uint16Slice_Int16Slice(i []uint16) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func Uint16Slice_Int16SlicePtr(i []uint16) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint16Slice_Int16Slice(i))
}
func Int16SlicePtr_Uint16Slice(i *[]int16, opt ...Option) []uint16 {
	return Int16Slice_Uint16Slice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Uint16SlicePtr(i *[]int16, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int16Slice_Uint16Slice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func Uint16SlicePtr_Int16Slice(i *[]uint16, opt ...Option) []int16 {
	return Uint16Slice_Int16Slice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Int16SlicePtr(i *[]uint16, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint16Slice_Int16Slice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}

func Int16_Uint32(i int16) uint32 {
	return uint32(i)
}
func Int16_Uint32Ptr(i int16) *uint32 {
	return Uint32_Uint32Ptr(Int16_Uint32(i))
}
func Uint32_Int16(i uint32) int16 {
	return int16(i)
}
func Uint32_Int16Ptr(i uint32) *int16 {
	return Int16_Int16Ptr(Uint32_Int16(i))
}
func Int16Ptr_Uint32(i *int16, opt ...Option) uint32 {
	return Int16_Uint32(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Uint32Ptr(i *int16, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Int16_Uint32(Int16Ptr_Int16(i, opt...)))
}
func Uint32Ptr_Int16(i *uint32, opt ...Option) int16 {
	return Uint32_Int16(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Int16Ptr(i *uint32, opt ...Option) *int16 {
	return Int16_Int16Ptr(Uint32_Int16(Uint32Ptr_Uint32(i, opt...)))
}

func Int16Slice_Uint32Slice(i []int16) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func Int16Slice_Uint32SlicePtr(i []int16) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int16Slice_Uint32Slice(i))
}
func Uint32Slice_Int16Slice(i []uint32) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func Uint32Slice_Int16SlicePtr(i []uint32) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint32Slice_Int16Slice(i))
}
func Int16SlicePtr_Uint32Slice(i *[]int16, opt ...Option) []uint32 {
	return Int16Slice_Uint32Slice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Uint32SlicePtr(i *[]int16, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int16Slice_Uint32Slice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func Uint32SlicePtr_Int16Slice(i *[]uint32, opt ...Option) []int16 {
	return Uint32Slice_Int16Slice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Int16SlicePtr(i *[]uint32, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint32Slice_Int16Slice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}

func Int16_Uint64(i int16) uint64 {
	return uint64(i)
}
func Int16_Uint64Ptr(i int16) *uint64 {
	return Uint64_Uint64Ptr(Int16_Uint64(i))
}
func Uint64_Int16(i uint64) int16 {
	return int16(i)
}
func Uint64_Int16Ptr(i uint64) *int16 {
	return Int16_Int16Ptr(Uint64_Int16(i))
}
func Int16Ptr_Uint64(i *int16, opt ...Option) uint64 {
	return Int16_Uint64(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Uint64Ptr(i *int16, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Int16_Uint64(Int16Ptr_Int16(i, opt...)))
}
func Uint64Ptr_Int16(i *uint64, opt ...Option) int16 {
	return Uint64_Int16(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Int16Ptr(i *uint64, opt ...Option) *int16 {
	return Int16_Int16Ptr(Uint64_Int16(Uint64Ptr_Uint64(i, opt...)))
}

func Int16Slice_Uint64Slice(i []int16) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func Int16Slice_Uint64SlicePtr(i []int16) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int16Slice_Uint64Slice(i))
}
func Uint64Slice_Int16Slice(i []uint64) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func Uint64Slice_Int16SlicePtr(i []uint64) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint64Slice_Int16Slice(i))
}
func Int16SlicePtr_Uint64Slice(i *[]int16, opt ...Option) []uint64 {
	return Int16Slice_Uint64Slice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Uint64SlicePtr(i *[]int16, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int16Slice_Uint64Slice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func Uint64SlicePtr_Int16Slice(i *[]uint64, opt ...Option) []int16 {
	return Uint64Slice_Int16Slice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Int16SlicePtr(i *[]uint64, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint64Slice_Int16Slice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Int16_Uint8(i int16) uint8 {
	return uint8(i)
}
func Int16_Uint8Ptr(i int16) *uint8 {
	return Uint8_Uint8Ptr(Int16_Uint8(i))
}
func Uint8_Int16(i uint8) int16 {
	return int16(i)
}
func Uint8_Int16Ptr(i uint8) *int16 {
	return Int16_Int16Ptr(Uint8_Int16(i))
}
func Int16Ptr_Uint8(i *int16, opt ...Option) uint8 {
	return Int16_Uint8(Int16Ptr_Int16(i, opt...))
}
func Int16Ptr_Uint8Ptr(i *int16, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Int16_Uint8(Int16Ptr_Int16(i, opt...)))
}
func Uint8Ptr_Int16(i *uint8, opt ...Option) int16 {
	return Uint8_Int16(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Int16Ptr(i *uint8, opt ...Option) *int16 {
	return Int16_Int16Ptr(Uint8_Int16(Uint8Ptr_Uint8(i, opt...)))
}

func Int16Slice_Uint8Slice(i []int16) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func Int16Slice_Uint8SlicePtr(i []int16) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int16Slice_Uint8Slice(i))
}
func Uint8Slice_Int16Slice(i []uint8) []int16 {
	res := []int16{}
	for _, item := range i {
		res = append(res, int16(item))
	}
	return res
}
func Uint8Slice_Int16SlicePtr(i []uint8) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint8Slice_Int16Slice(i))
}
func Int16SlicePtr_Uint8Slice(i *[]int16, opt ...Option) []uint8 {
	return Int16Slice_Uint8Slice(Int16SlicePtr_Int16Slice(i, opt...))
}
func Int16SlicePtr_Uint8SlicePtr(i *[]int16, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int16Slice_Uint8Slice(Int16SlicePtr_Int16Slice(i, opt...)))
}
func Uint8SlicePtr_Int16Slice(i *[]uint8, opt ...Option) []int16 {
	return Uint8Slice_Int16Slice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Int16SlicePtr(i *[]uint8, opt ...Option) *[]int16 {
	return Int16Slice_Int16SlicePtr(Uint8Slice_Int16Slice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Int32_Int64(i int32) int64 {
	return int64(i)
}
func Int32_Int64Ptr(i int32) *int64 {
	return Int64_Int64Ptr(Int32_Int64(i))
}
func Int64_Int32(i int64) int32 {
	return int32(i)
}
func Int64_Int32Ptr(i int64) *int32 {
	return Int32_Int32Ptr(Int64_Int32(i))
}
func Int32Ptr_Int64(i *int32, opt ...Option) int64 {
	return Int32_Int64(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Int64Ptr(i *int32, opt ...Option) *int64 {
	return Int64_Int64Ptr(Int32_Int64(Int32Ptr_Int32(i, opt...)))
}
func Int64Ptr_Int32(i *int64, opt ...Option) int32 {
	return Int64_Int32(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Int32Ptr(i *int64, opt ...Option) *int32 {
	return Int32_Int32Ptr(Int64_Int32(Int64Ptr_Int64(i, opt...)))
}

func Int32Slice_Int64Slice(i []int32) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func Int32Slice_Int64SlicePtr(i []int32) *[]int64 {
	return Int64Slice_Int64SlicePtr(Int32Slice_Int64Slice(i))
}
func Int64Slice_Int32Slice(i []int64) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func Int64Slice_Int32SlicePtr(i []int64) *[]int32 {
	return Int32Slice_Int32SlicePtr(Int64Slice_Int32Slice(i))
}
func Int32SlicePtr_Int64Slice(i *[]int32, opt ...Option) []int64 {
	return Int32Slice_Int64Slice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Int64SlicePtr(i *[]int32, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(Int32Slice_Int64Slice(Int32SlicePtr_Int32Slice(i, opt...)))
}
func Int64SlicePtr_Int32Slice(i *[]int64, opt ...Option) []int32 {
	return Int64Slice_Int32Slice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Int32SlicePtr(i *[]int64, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(Int64Slice_Int32Slice(Int64SlicePtr_Int64Slice(i, opt...)))
}

func Int32_Int8(i int32) int8 {
	return int8(i)
}
func Int32_Int8Ptr(i int32) *int8 {
	return Int8_Int8Ptr(Int32_Int8(i))
}
func Int8_Int32(i int8) int32 {
	return int32(i)
}
func Int8_Int32Ptr(i int8) *int32 {
	return Int32_Int32Ptr(Int8_Int32(i))
}
func Int32Ptr_Int8(i *int32, opt ...Option) int8 {
	return Int32_Int8(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Int8Ptr(i *int32, opt ...Option) *int8 {
	return Int8_Int8Ptr(Int32_Int8(Int32Ptr_Int32(i, opt...)))
}
func Int8Ptr_Int32(i *int8, opt ...Option) int32 {
	return Int8_Int32(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Int32Ptr(i *int8, opt ...Option) *int32 {
	return Int32_Int32Ptr(Int8_Int32(Int8Ptr_Int8(i, opt...)))
}

func Int32Slice_Int8Slice(i []int32) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func Int32Slice_Int8SlicePtr(i []int32) *[]int8 {
	return Int8Slice_Int8SlicePtr(Int32Slice_Int8Slice(i))
}
func Int8Slice_Int32Slice(i []int8) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func Int8Slice_Int32SlicePtr(i []int8) *[]int32 {
	return Int32Slice_Int32SlicePtr(Int8Slice_Int32Slice(i))
}
func Int32SlicePtr_Int8Slice(i *[]int32, opt ...Option) []int8 {
	return Int32Slice_Int8Slice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Int8SlicePtr(i *[]int32, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(Int32Slice_Int8Slice(Int32SlicePtr_Int32Slice(i, opt...)))
}
func Int8SlicePtr_Int32Slice(i *[]int8, opt ...Option) []int32 {
	return Int8Slice_Int32Slice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Int32SlicePtr(i *[]int8, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(Int8Slice_Int32Slice(Int8SlicePtr_Int8Slice(i, opt...)))
}

func Int32_Uint(i int32) uint {
	return uint(i)
}
func Int32_UintPtr(i int32) *uint {
	return Uint_UintPtr(Int32_Uint(i))
}
func Uint_Int32(i uint) int32 {
	return int32(i)
}
func Uint_Int32Ptr(i uint) *int32 {
	return Int32_Int32Ptr(Uint_Int32(i))
}
func Int32Ptr_Uint(i *int32, opt ...Option) uint {
	return Int32_Uint(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_UintPtr(i *int32, opt ...Option) *uint {
	return Uint_UintPtr(Int32_Uint(Int32Ptr_Int32(i, opt...)))
}
func UintPtr_Int32(i *uint, opt ...Option) int32 {
	return Uint_Int32(UintPtr_Uint(i, opt...))
}
func UintPtr_Int32Ptr(i *uint, opt ...Option) *int32 {
	return Int32_Int32Ptr(Uint_Int32(UintPtr_Uint(i, opt...)))
}

func Int32Slice_UintSlice(i []int32) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Int32Slice_UintSlicePtr(i []int32) *[]uint {
	return UintSlice_UintSlicePtr(Int32Slice_UintSlice(i))
}
func UintSlice_Int32Slice(i []uint) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func UintSlice_Int32SlicePtr(i []uint) *[]int32 {
	return Int32Slice_Int32SlicePtr(UintSlice_Int32Slice(i))
}
func Int32SlicePtr_UintSlice(i *[]int32, opt ...Option) []uint {
	return Int32Slice_UintSlice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_UintSlicePtr(i *[]int32, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Int32Slice_UintSlice(Int32SlicePtr_Int32Slice(i, opt...)))
}
func UintSlicePtr_Int32Slice(i *[]uint, opt ...Option) []int32 {
	return UintSlice_Int32Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Int32SlicePtr(i *[]uint, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(UintSlice_Int32Slice(UintSlicePtr_UintSlice(i, opt...)))
}

func Int32_Uint16(i int32) uint16 {
	return uint16(i)
}
func Int32_Uint16Ptr(i int32) *uint16 {
	return Uint16_Uint16Ptr(Int32_Uint16(i))
}
func Uint16_Int32(i uint16) int32 {
	return int32(i)
}
func Uint16_Int32Ptr(i uint16) *int32 {
	return Int32_Int32Ptr(Uint16_Int32(i))
}
func Int32Ptr_Uint16(i *int32, opt ...Option) uint16 {
	return Int32_Uint16(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Uint16Ptr(i *int32, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Int32_Uint16(Int32Ptr_Int32(i, opt...)))
}
func Uint16Ptr_Int32(i *uint16, opt ...Option) int32 {
	return Uint16_Int32(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Int32Ptr(i *uint16, opt ...Option) *int32 {
	return Int32_Int32Ptr(Uint16_Int32(Uint16Ptr_Uint16(i, opt...)))
}

func Int32Slice_Uint16Slice(i []int32) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func Int32Slice_Uint16SlicePtr(i []int32) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int32Slice_Uint16Slice(i))
}
func Uint16Slice_Int32Slice(i []uint16) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func Uint16Slice_Int32SlicePtr(i []uint16) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint16Slice_Int32Slice(i))
}
func Int32SlicePtr_Uint16Slice(i *[]int32, opt ...Option) []uint16 {
	return Int32Slice_Uint16Slice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Uint16SlicePtr(i *[]int32, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int32Slice_Uint16Slice(Int32SlicePtr_Int32Slice(i, opt...)))
}
func Uint16SlicePtr_Int32Slice(i *[]uint16, opt ...Option) []int32 {
	return Uint16Slice_Int32Slice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Int32SlicePtr(i *[]uint16, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint16Slice_Int32Slice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}

func Int32_Uint32(i int32) uint32 {
	return uint32(i)
}
func Int32_Uint32Ptr(i int32) *uint32 {
	return Uint32_Uint32Ptr(Int32_Uint32(i))
}
func Uint32_Int32(i uint32) int32 {
	return int32(i)
}
func Uint32_Int32Ptr(i uint32) *int32 {
	return Int32_Int32Ptr(Uint32_Int32(i))
}
func Int32Ptr_Uint32(i *int32, opt ...Option) uint32 {
	return Int32_Uint32(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Uint32Ptr(i *int32, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Int32_Uint32(Int32Ptr_Int32(i, opt...)))
}
func Uint32Ptr_Int32(i *uint32, opt ...Option) int32 {
	return Uint32_Int32(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Int32Ptr(i *uint32, opt ...Option) *int32 {
	return Int32_Int32Ptr(Uint32_Int32(Uint32Ptr_Uint32(i, opt...)))
}

func Int32Slice_Uint32Slice(i []int32) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func Int32Slice_Uint32SlicePtr(i []int32) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int32Slice_Uint32Slice(i))
}
func Uint32Slice_Int32Slice(i []uint32) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func Uint32Slice_Int32SlicePtr(i []uint32) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint32Slice_Int32Slice(i))
}
func Int32SlicePtr_Uint32Slice(i *[]int32, opt ...Option) []uint32 {
	return Int32Slice_Uint32Slice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Uint32SlicePtr(i *[]int32, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int32Slice_Uint32Slice(Int32SlicePtr_Int32Slice(i, opt...)))
}
func Uint32SlicePtr_Int32Slice(i *[]uint32, opt ...Option) []int32 {
	return Uint32Slice_Int32Slice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Int32SlicePtr(i *[]uint32, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint32Slice_Int32Slice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}

func Int32_Uint64(i int32) uint64 {
	return uint64(i)
}
func Int32_Uint64Ptr(i int32) *uint64 {
	return Uint64_Uint64Ptr(Int32_Uint64(i))
}
func Uint64_Int32(i uint64) int32 {
	return int32(i)
}
func Uint64_Int32Ptr(i uint64) *int32 {
	return Int32_Int32Ptr(Uint64_Int32(i))
}
func Int32Ptr_Uint64(i *int32, opt ...Option) uint64 {
	return Int32_Uint64(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Uint64Ptr(i *int32, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Int32_Uint64(Int32Ptr_Int32(i, opt...)))
}
func Uint64Ptr_Int32(i *uint64, opt ...Option) int32 {
	return Uint64_Int32(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Int32Ptr(i *uint64, opt ...Option) *int32 {
	return Int32_Int32Ptr(Uint64_Int32(Uint64Ptr_Uint64(i, opt...)))
}

func Int32Slice_Uint64Slice(i []int32) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func Int32Slice_Uint64SlicePtr(i []int32) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int32Slice_Uint64Slice(i))
}
func Uint64Slice_Int32Slice(i []uint64) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func Uint64Slice_Int32SlicePtr(i []uint64) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint64Slice_Int32Slice(i))
}
func Int32SlicePtr_Uint64Slice(i *[]int32, opt ...Option) []uint64 {
	return Int32Slice_Uint64Slice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Uint64SlicePtr(i *[]int32, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int32Slice_Uint64Slice(Int32SlicePtr_Int32Slice(i, opt...)))
}
func Uint64SlicePtr_Int32Slice(i *[]uint64, opt ...Option) []int32 {
	return Uint64Slice_Int32Slice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Int32SlicePtr(i *[]uint64, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint64Slice_Int32Slice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Int32_Uint8(i int32) uint8 {
	return uint8(i)
}
func Int32_Uint8Ptr(i int32) *uint8 {
	return Uint8_Uint8Ptr(Int32_Uint8(i))
}
func Uint8_Int32(i uint8) int32 {
	return int32(i)
}
func Uint8_Int32Ptr(i uint8) *int32 {
	return Int32_Int32Ptr(Uint8_Int32(i))
}
func Int32Ptr_Uint8(i *int32, opt ...Option) uint8 {
	return Int32_Uint8(Int32Ptr_Int32(i, opt...))
}
func Int32Ptr_Uint8Ptr(i *int32, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Int32_Uint8(Int32Ptr_Int32(i, opt...)))
}
func Uint8Ptr_Int32(i *uint8, opt ...Option) int32 {
	return Uint8_Int32(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Int32Ptr(i *uint8, opt ...Option) *int32 {
	return Int32_Int32Ptr(Uint8_Int32(Uint8Ptr_Uint8(i, opt...)))
}

func Int32Slice_Uint8Slice(i []int32) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func Int32Slice_Uint8SlicePtr(i []int32) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int32Slice_Uint8Slice(i))
}
func Uint8Slice_Int32Slice(i []uint8) []int32 {
	res := []int32{}
	for _, item := range i {
		res = append(res, int32(item))
	}
	return res
}
func Uint8Slice_Int32SlicePtr(i []uint8) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint8Slice_Int32Slice(i))
}
func Int32SlicePtr_Uint8Slice(i *[]int32, opt ...Option) []uint8 {
	return Int32Slice_Uint8Slice(Int32SlicePtr_Int32Slice(i, opt...))
}
func Int32SlicePtr_Uint8SlicePtr(i *[]int32, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int32Slice_Uint8Slice(Int32SlicePtr_Int32Slice(i, opt...)))
}
func Uint8SlicePtr_Int32Slice(i *[]uint8, opt ...Option) []int32 {
	return Uint8Slice_Int32Slice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Int32SlicePtr(i *[]uint8, opt ...Option) *[]int32 {
	return Int32Slice_Int32SlicePtr(Uint8Slice_Int32Slice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Int64_Int8(i int64) int8 {
	return int8(i)
}
func Int64_Int8Ptr(i int64) *int8 {
	return Int8_Int8Ptr(Int64_Int8(i))
}
func Int8_Int64(i int8) int64 {
	return int64(i)
}
func Int8_Int64Ptr(i int8) *int64 {
	return Int64_Int64Ptr(Int8_Int64(i))
}
func Int64Ptr_Int8(i *int64, opt ...Option) int8 {
	return Int64_Int8(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Int8Ptr(i *int64, opt ...Option) *int8 {
	return Int8_Int8Ptr(Int64_Int8(Int64Ptr_Int64(i, opt...)))
}
func Int8Ptr_Int64(i *int8, opt ...Option) int64 {
	return Int8_Int64(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Int64Ptr(i *int8, opt ...Option) *int64 {
	return Int64_Int64Ptr(Int8_Int64(Int8Ptr_Int8(i, opt...)))
}

func Int64Slice_Int8Slice(i []int64) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func Int64Slice_Int8SlicePtr(i []int64) *[]int8 {
	return Int8Slice_Int8SlicePtr(Int64Slice_Int8Slice(i))
}
func Int8Slice_Int64Slice(i []int8) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func Int8Slice_Int64SlicePtr(i []int8) *[]int64 {
	return Int64Slice_Int64SlicePtr(Int8Slice_Int64Slice(i))
}
func Int64SlicePtr_Int8Slice(i *[]int64, opt ...Option) []int8 {
	return Int64Slice_Int8Slice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Int8SlicePtr(i *[]int64, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(Int64Slice_Int8Slice(Int64SlicePtr_Int64Slice(i, opt...)))
}
func Int8SlicePtr_Int64Slice(i *[]int8, opt ...Option) []int64 {
	return Int8Slice_Int64Slice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Int64SlicePtr(i *[]int8, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(Int8Slice_Int64Slice(Int8SlicePtr_Int8Slice(i, opt...)))
}

func Int64_Uint(i int64) uint {
	return uint(i)
}
func Int64_UintPtr(i int64) *uint {
	return Uint_UintPtr(Int64_Uint(i))
}
func Uint_Int64(i uint) int64 {
	return int64(i)
}
func Uint_Int64Ptr(i uint) *int64 {
	return Int64_Int64Ptr(Uint_Int64(i))
}
func Int64Ptr_Uint(i *int64, opt ...Option) uint {
	return Int64_Uint(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_UintPtr(i *int64, opt ...Option) *uint {
	return Uint_UintPtr(Int64_Uint(Int64Ptr_Int64(i, opt...)))
}
func UintPtr_Int64(i *uint, opt ...Option) int64 {
	return Uint_Int64(UintPtr_Uint(i, opt...))
}
func UintPtr_Int64Ptr(i *uint, opt ...Option) *int64 {
	return Int64_Int64Ptr(Uint_Int64(UintPtr_Uint(i, opt...)))
}

func Int64Slice_UintSlice(i []int64) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Int64Slice_UintSlicePtr(i []int64) *[]uint {
	return UintSlice_UintSlicePtr(Int64Slice_UintSlice(i))
}
func UintSlice_Int64Slice(i []uint) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func UintSlice_Int64SlicePtr(i []uint) *[]int64 {
	return Int64Slice_Int64SlicePtr(UintSlice_Int64Slice(i))
}
func Int64SlicePtr_UintSlice(i *[]int64, opt ...Option) []uint {
	return Int64Slice_UintSlice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_UintSlicePtr(i *[]int64, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Int64Slice_UintSlice(Int64SlicePtr_Int64Slice(i, opt...)))
}
func UintSlicePtr_Int64Slice(i *[]uint, opt ...Option) []int64 {
	return UintSlice_Int64Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Int64SlicePtr(i *[]uint, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(UintSlice_Int64Slice(UintSlicePtr_UintSlice(i, opt...)))
}

func Int64_Uint16(i int64) uint16 {
	return uint16(i)
}
func Int64_Uint16Ptr(i int64) *uint16 {
	return Uint16_Uint16Ptr(Int64_Uint16(i))
}
func Uint16_Int64(i uint16) int64 {
	return int64(i)
}
func Uint16_Int64Ptr(i uint16) *int64 {
	return Int64_Int64Ptr(Uint16_Int64(i))
}
func Int64Ptr_Uint16(i *int64, opt ...Option) uint16 {
	return Int64_Uint16(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Uint16Ptr(i *int64, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Int64_Uint16(Int64Ptr_Int64(i, opt...)))
}
func Uint16Ptr_Int64(i *uint16, opt ...Option) int64 {
	return Uint16_Int64(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Int64Ptr(i *uint16, opt ...Option) *int64 {
	return Int64_Int64Ptr(Uint16_Int64(Uint16Ptr_Uint16(i, opt...)))
}

func Int64Slice_Uint16Slice(i []int64) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func Int64Slice_Uint16SlicePtr(i []int64) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int64Slice_Uint16Slice(i))
}
func Uint16Slice_Int64Slice(i []uint16) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func Uint16Slice_Int64SlicePtr(i []uint16) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint16Slice_Int64Slice(i))
}
func Int64SlicePtr_Uint16Slice(i *[]int64, opt ...Option) []uint16 {
	return Int64Slice_Uint16Slice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Uint16SlicePtr(i *[]int64, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int64Slice_Uint16Slice(Int64SlicePtr_Int64Slice(i, opt...)))
}
func Uint16SlicePtr_Int64Slice(i *[]uint16, opt ...Option) []int64 {
	return Uint16Slice_Int64Slice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Int64SlicePtr(i *[]uint16, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint16Slice_Int64Slice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}

func Int64_Uint32(i int64) uint32 {
	return uint32(i)
}
func Int64_Uint32Ptr(i int64) *uint32 {
	return Uint32_Uint32Ptr(Int64_Uint32(i))
}
func Uint32_Int64(i uint32) int64 {
	return int64(i)
}
func Uint32_Int64Ptr(i uint32) *int64 {
	return Int64_Int64Ptr(Uint32_Int64(i))
}
func Int64Ptr_Uint32(i *int64, opt ...Option) uint32 {
	return Int64_Uint32(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Uint32Ptr(i *int64, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Int64_Uint32(Int64Ptr_Int64(i, opt...)))
}
func Uint32Ptr_Int64(i *uint32, opt ...Option) int64 {
	return Uint32_Int64(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Int64Ptr(i *uint32, opt ...Option) *int64 {
	return Int64_Int64Ptr(Uint32_Int64(Uint32Ptr_Uint32(i, opt...)))
}

func Int64Slice_Uint32Slice(i []int64) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func Int64Slice_Uint32SlicePtr(i []int64) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int64Slice_Uint32Slice(i))
}
func Uint32Slice_Int64Slice(i []uint32) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func Uint32Slice_Int64SlicePtr(i []uint32) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint32Slice_Int64Slice(i))
}
func Int64SlicePtr_Uint32Slice(i *[]int64, opt ...Option) []uint32 {
	return Int64Slice_Uint32Slice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Uint32SlicePtr(i *[]int64, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int64Slice_Uint32Slice(Int64SlicePtr_Int64Slice(i, opt...)))
}
func Uint32SlicePtr_Int64Slice(i *[]uint32, opt ...Option) []int64 {
	return Uint32Slice_Int64Slice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Int64SlicePtr(i *[]uint32, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint32Slice_Int64Slice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}

func Int64_Uint64(i int64) uint64 {
	return uint64(i)
}
func Int64_Uint64Ptr(i int64) *uint64 {
	return Uint64_Uint64Ptr(Int64_Uint64(i))
}
func Uint64_Int64(i uint64) int64 {
	return int64(i)
}
func Uint64_Int64Ptr(i uint64) *int64 {
	return Int64_Int64Ptr(Uint64_Int64(i))
}
func Int64Ptr_Uint64(i *int64, opt ...Option) uint64 {
	return Int64_Uint64(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Uint64Ptr(i *int64, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Int64_Uint64(Int64Ptr_Int64(i, opt...)))
}
func Uint64Ptr_Int64(i *uint64, opt ...Option) int64 {
	return Uint64_Int64(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Int64Ptr(i *uint64, opt ...Option) *int64 {
	return Int64_Int64Ptr(Uint64_Int64(Uint64Ptr_Uint64(i, opt...)))
}

func Int64Slice_Uint64Slice(i []int64) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func Int64Slice_Uint64SlicePtr(i []int64) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int64Slice_Uint64Slice(i))
}
func Uint64Slice_Int64Slice(i []uint64) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func Uint64Slice_Int64SlicePtr(i []uint64) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint64Slice_Int64Slice(i))
}
func Int64SlicePtr_Uint64Slice(i *[]int64, opt ...Option) []uint64 {
	return Int64Slice_Uint64Slice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Uint64SlicePtr(i *[]int64, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int64Slice_Uint64Slice(Int64SlicePtr_Int64Slice(i, opt...)))
}
func Uint64SlicePtr_Int64Slice(i *[]uint64, opt ...Option) []int64 {
	return Uint64Slice_Int64Slice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Int64SlicePtr(i *[]uint64, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint64Slice_Int64Slice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Int64_Uint8(i int64) uint8 {
	return uint8(i)
}
func Int64_Uint8Ptr(i int64) *uint8 {
	return Uint8_Uint8Ptr(Int64_Uint8(i))
}
func Uint8_Int64(i uint8) int64 {
	return int64(i)
}
func Uint8_Int64Ptr(i uint8) *int64 {
	return Int64_Int64Ptr(Uint8_Int64(i))
}
func Int64Ptr_Uint8(i *int64, opt ...Option) uint8 {
	return Int64_Uint8(Int64Ptr_Int64(i, opt...))
}
func Int64Ptr_Uint8Ptr(i *int64, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Int64_Uint8(Int64Ptr_Int64(i, opt...)))
}
func Uint8Ptr_Int64(i *uint8, opt ...Option) int64 {
	return Uint8_Int64(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Int64Ptr(i *uint8, opt ...Option) *int64 {
	return Int64_Int64Ptr(Uint8_Int64(Uint8Ptr_Uint8(i, opt...)))
}

func Int64Slice_Uint8Slice(i []int64) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func Int64Slice_Uint8SlicePtr(i []int64) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int64Slice_Uint8Slice(i))
}
func Uint8Slice_Int64Slice(i []uint8) []int64 {
	res := []int64{}
	for _, item := range i {
		res = append(res, int64(item))
	}
	return res
}
func Uint8Slice_Int64SlicePtr(i []uint8) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint8Slice_Int64Slice(i))
}
func Int64SlicePtr_Uint8Slice(i *[]int64, opt ...Option) []uint8 {
	return Int64Slice_Uint8Slice(Int64SlicePtr_Int64Slice(i, opt...))
}
func Int64SlicePtr_Uint8SlicePtr(i *[]int64, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int64Slice_Uint8Slice(Int64SlicePtr_Int64Slice(i, opt...)))
}
func Uint8SlicePtr_Int64Slice(i *[]uint8, opt ...Option) []int64 {
	return Uint8Slice_Int64Slice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Int64SlicePtr(i *[]uint8, opt ...Option) *[]int64 {
	return Int64Slice_Int64SlicePtr(Uint8Slice_Int64Slice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Int8_Uint(i int8) uint {
	return uint(i)
}
func Int8_UintPtr(i int8) *uint {
	return Uint_UintPtr(Int8_Uint(i))
}
func Uint_Int8(i uint) int8 {
	return int8(i)
}
func Uint_Int8Ptr(i uint) *int8 {
	return Int8_Int8Ptr(Uint_Int8(i))
}
func Int8Ptr_Uint(i *int8, opt ...Option) uint {
	return Int8_Uint(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_UintPtr(i *int8, opt ...Option) *uint {
	return Uint_UintPtr(Int8_Uint(Int8Ptr_Int8(i, opt...)))
}
func UintPtr_Int8(i *uint, opt ...Option) int8 {
	return Uint_Int8(UintPtr_Uint(i, opt...))
}
func UintPtr_Int8Ptr(i *uint, opt ...Option) *int8 {
	return Int8_Int8Ptr(Uint_Int8(UintPtr_Uint(i, opt...)))
}

func Int8Slice_UintSlice(i []int8) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Int8Slice_UintSlicePtr(i []int8) *[]uint {
	return UintSlice_UintSlicePtr(Int8Slice_UintSlice(i))
}
func UintSlice_Int8Slice(i []uint) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func UintSlice_Int8SlicePtr(i []uint) *[]int8 {
	return Int8Slice_Int8SlicePtr(UintSlice_Int8Slice(i))
}
func Int8SlicePtr_UintSlice(i *[]int8, opt ...Option) []uint {
	return Int8Slice_UintSlice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_UintSlicePtr(i *[]int8, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Int8Slice_UintSlice(Int8SlicePtr_Int8Slice(i, opt...)))
}
func UintSlicePtr_Int8Slice(i *[]uint, opt ...Option) []int8 {
	return UintSlice_Int8Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Int8SlicePtr(i *[]uint, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(UintSlice_Int8Slice(UintSlicePtr_UintSlice(i, opt...)))
}

func Int8_Uint16(i int8) uint16 {
	return uint16(i)
}
func Int8_Uint16Ptr(i int8) *uint16 {
	return Uint16_Uint16Ptr(Int8_Uint16(i))
}
func Uint16_Int8(i uint16) int8 {
	return int8(i)
}
func Uint16_Int8Ptr(i uint16) *int8 {
	return Int8_Int8Ptr(Uint16_Int8(i))
}
func Int8Ptr_Uint16(i *int8, opt ...Option) uint16 {
	return Int8_Uint16(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Uint16Ptr(i *int8, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Int8_Uint16(Int8Ptr_Int8(i, opt...)))
}
func Uint16Ptr_Int8(i *uint16, opt ...Option) int8 {
	return Uint16_Int8(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Int8Ptr(i *uint16, opt ...Option) *int8 {
	return Int8_Int8Ptr(Uint16_Int8(Uint16Ptr_Uint16(i, opt...)))
}

func Int8Slice_Uint16Slice(i []int8) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func Int8Slice_Uint16SlicePtr(i []int8) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int8Slice_Uint16Slice(i))
}
func Uint16Slice_Int8Slice(i []uint16) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func Uint16Slice_Int8SlicePtr(i []uint16) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint16Slice_Int8Slice(i))
}
func Int8SlicePtr_Uint16Slice(i *[]int8, opt ...Option) []uint16 {
	return Int8Slice_Uint16Slice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Uint16SlicePtr(i *[]int8, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Int8Slice_Uint16Slice(Int8SlicePtr_Int8Slice(i, opt...)))
}
func Uint16SlicePtr_Int8Slice(i *[]uint16, opt ...Option) []int8 {
	return Uint16Slice_Int8Slice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Int8SlicePtr(i *[]uint16, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint16Slice_Int8Slice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}

func Int8_Uint32(i int8) uint32 {
	return uint32(i)
}
func Int8_Uint32Ptr(i int8) *uint32 {
	return Uint32_Uint32Ptr(Int8_Uint32(i))
}
func Uint32_Int8(i uint32) int8 {
	return int8(i)
}
func Uint32_Int8Ptr(i uint32) *int8 {
	return Int8_Int8Ptr(Uint32_Int8(i))
}
func Int8Ptr_Uint32(i *int8, opt ...Option) uint32 {
	return Int8_Uint32(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Uint32Ptr(i *int8, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Int8_Uint32(Int8Ptr_Int8(i, opt...)))
}
func Uint32Ptr_Int8(i *uint32, opt ...Option) int8 {
	return Uint32_Int8(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Int8Ptr(i *uint32, opt ...Option) *int8 {
	return Int8_Int8Ptr(Uint32_Int8(Uint32Ptr_Uint32(i, opt...)))
}

func Int8Slice_Uint32Slice(i []int8) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func Int8Slice_Uint32SlicePtr(i []int8) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int8Slice_Uint32Slice(i))
}
func Uint32Slice_Int8Slice(i []uint32) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func Uint32Slice_Int8SlicePtr(i []uint32) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint32Slice_Int8Slice(i))
}
func Int8SlicePtr_Uint32Slice(i *[]int8, opt ...Option) []uint32 {
	return Int8Slice_Uint32Slice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Uint32SlicePtr(i *[]int8, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Int8Slice_Uint32Slice(Int8SlicePtr_Int8Slice(i, opt...)))
}
func Uint32SlicePtr_Int8Slice(i *[]uint32, opt ...Option) []int8 {
	return Uint32Slice_Int8Slice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Int8SlicePtr(i *[]uint32, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint32Slice_Int8Slice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}

func Int8_Uint64(i int8) uint64 {
	return uint64(i)
}
func Int8_Uint64Ptr(i int8) *uint64 {
	return Uint64_Uint64Ptr(Int8_Uint64(i))
}
func Uint64_Int8(i uint64) int8 {
	return int8(i)
}
func Uint64_Int8Ptr(i uint64) *int8 {
	return Int8_Int8Ptr(Uint64_Int8(i))
}
func Int8Ptr_Uint64(i *int8, opt ...Option) uint64 {
	return Int8_Uint64(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Uint64Ptr(i *int8, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Int8_Uint64(Int8Ptr_Int8(i, opt...)))
}
func Uint64Ptr_Int8(i *uint64, opt ...Option) int8 {
	return Uint64_Int8(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Int8Ptr(i *uint64, opt ...Option) *int8 {
	return Int8_Int8Ptr(Uint64_Int8(Uint64Ptr_Uint64(i, opt...)))
}

func Int8Slice_Uint64Slice(i []int8) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func Int8Slice_Uint64SlicePtr(i []int8) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int8Slice_Uint64Slice(i))
}
func Uint64Slice_Int8Slice(i []uint64) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func Uint64Slice_Int8SlicePtr(i []uint64) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint64Slice_Int8Slice(i))
}
func Int8SlicePtr_Uint64Slice(i *[]int8, opt ...Option) []uint64 {
	return Int8Slice_Uint64Slice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Uint64SlicePtr(i *[]int8, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Int8Slice_Uint64Slice(Int8SlicePtr_Int8Slice(i, opt...)))
}
func Uint64SlicePtr_Int8Slice(i *[]uint64, opt ...Option) []int8 {
	return Uint64Slice_Int8Slice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Int8SlicePtr(i *[]uint64, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint64Slice_Int8Slice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Int8_Uint8(i int8) uint8 {
	return uint8(i)
}
func Int8_Uint8Ptr(i int8) *uint8 {
	return Uint8_Uint8Ptr(Int8_Uint8(i))
}
func Uint8_Int8(i uint8) int8 {
	return int8(i)
}
func Uint8_Int8Ptr(i uint8) *int8 {
	return Int8_Int8Ptr(Uint8_Int8(i))
}
func Int8Ptr_Uint8(i *int8, opt ...Option) uint8 {
	return Int8_Uint8(Int8Ptr_Int8(i, opt...))
}
func Int8Ptr_Uint8Ptr(i *int8, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Int8_Uint8(Int8Ptr_Int8(i, opt...)))
}
func Uint8Ptr_Int8(i *uint8, opt ...Option) int8 {
	return Uint8_Int8(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Int8Ptr(i *uint8, opt ...Option) *int8 {
	return Int8_Int8Ptr(Uint8_Int8(Uint8Ptr_Uint8(i, opt...)))
}

func Int8Slice_Uint8Slice(i []int8) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func Int8Slice_Uint8SlicePtr(i []int8) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int8Slice_Uint8Slice(i))
}
func Uint8Slice_Int8Slice(i []uint8) []int8 {
	res := []int8{}
	for _, item := range i {
		res = append(res, int8(item))
	}
	return res
}
func Uint8Slice_Int8SlicePtr(i []uint8) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint8Slice_Int8Slice(i))
}
func Int8SlicePtr_Uint8Slice(i *[]int8, opt ...Option) []uint8 {
	return Int8Slice_Uint8Slice(Int8SlicePtr_Int8Slice(i, opt...))
}
func Int8SlicePtr_Uint8SlicePtr(i *[]int8, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Int8Slice_Uint8Slice(Int8SlicePtr_Int8Slice(i, opt...)))
}
func Uint8SlicePtr_Int8Slice(i *[]uint8, opt ...Option) []int8 {
	return Uint8Slice_Int8Slice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Int8SlicePtr(i *[]uint8, opt ...Option) *[]int8 {
	return Int8Slice_Int8SlicePtr(Uint8Slice_Int8Slice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Uint_Uint16(i uint) uint16 {
	return uint16(i)
}
func Uint_Uint16Ptr(i uint) *uint16 {
	return Uint16_Uint16Ptr(Uint_Uint16(i))
}
func Uint16_Uint(i uint16) uint {
	return uint(i)
}
func Uint16_UintPtr(i uint16) *uint {
	return Uint_UintPtr(Uint16_Uint(i))
}
func UintPtr_Uint16(i *uint, opt ...Option) uint16 {
	return Uint_Uint16(UintPtr_Uint(i, opt...))
}
func UintPtr_Uint16Ptr(i *uint, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Uint_Uint16(UintPtr_Uint(i, opt...)))
}
func Uint16Ptr_Uint(i *uint16, opt ...Option) uint {
	return Uint16_Uint(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_UintPtr(i *uint16, opt ...Option) *uint {
	return Uint_UintPtr(Uint16_Uint(Uint16Ptr_Uint16(i, opt...)))
}

func UintSlice_Uint16Slice(i []uint) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func UintSlice_Uint16SlicePtr(i []uint) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(UintSlice_Uint16Slice(i))
}
func Uint16Slice_UintSlice(i []uint16) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Uint16Slice_UintSlicePtr(i []uint16) *[]uint {
	return UintSlice_UintSlicePtr(Uint16Slice_UintSlice(i))
}
func UintSlicePtr_Uint16Slice(i *[]uint, opt ...Option) []uint16 {
	return UintSlice_Uint16Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Uint16SlicePtr(i *[]uint, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(UintSlice_Uint16Slice(UintSlicePtr_UintSlice(i, opt...)))
}
func Uint16SlicePtr_UintSlice(i *[]uint16, opt ...Option) []uint {
	return Uint16Slice_UintSlice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_UintSlicePtr(i *[]uint16, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Uint16Slice_UintSlice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}

func Uint_Uint32(i uint) uint32 {
	return uint32(i)
}
func Uint_Uint32Ptr(i uint) *uint32 {
	return Uint32_Uint32Ptr(Uint_Uint32(i))
}
func Uint32_Uint(i uint32) uint {
	return uint(i)
}
func Uint32_UintPtr(i uint32) *uint {
	return Uint_UintPtr(Uint32_Uint(i))
}
func UintPtr_Uint32(i *uint, opt ...Option) uint32 {
	return Uint_Uint32(UintPtr_Uint(i, opt...))
}
func UintPtr_Uint32Ptr(i *uint, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Uint_Uint32(UintPtr_Uint(i, opt...)))
}
func Uint32Ptr_Uint(i *uint32, opt ...Option) uint {
	return Uint32_Uint(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_UintPtr(i *uint32, opt ...Option) *uint {
	return Uint_UintPtr(Uint32_Uint(Uint32Ptr_Uint32(i, opt...)))
}

func UintSlice_Uint32Slice(i []uint) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func UintSlice_Uint32SlicePtr(i []uint) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(UintSlice_Uint32Slice(i))
}
func Uint32Slice_UintSlice(i []uint32) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Uint32Slice_UintSlicePtr(i []uint32) *[]uint {
	return UintSlice_UintSlicePtr(Uint32Slice_UintSlice(i))
}
func UintSlicePtr_Uint32Slice(i *[]uint, opt ...Option) []uint32 {
	return UintSlice_Uint32Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Uint32SlicePtr(i *[]uint, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(UintSlice_Uint32Slice(UintSlicePtr_UintSlice(i, opt...)))
}
func Uint32SlicePtr_UintSlice(i *[]uint32, opt ...Option) []uint {
	return Uint32Slice_UintSlice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_UintSlicePtr(i *[]uint32, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Uint32Slice_UintSlice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}

func Uint_Uint64(i uint) uint64 {
	return uint64(i)
}
func Uint_Uint64Ptr(i uint) *uint64 {
	return Uint64_Uint64Ptr(Uint_Uint64(i))
}
func Uint64_Uint(i uint64) uint {
	return uint(i)
}
func Uint64_UintPtr(i uint64) *uint {
	return Uint_UintPtr(Uint64_Uint(i))
}
func UintPtr_Uint64(i *uint, opt ...Option) uint64 {
	return Uint_Uint64(UintPtr_Uint(i, opt...))
}
func UintPtr_Uint64Ptr(i *uint, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Uint_Uint64(UintPtr_Uint(i, opt...)))
}
func Uint64Ptr_Uint(i *uint64, opt ...Option) uint {
	return Uint64_Uint(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_UintPtr(i *uint64, opt ...Option) *uint {
	return Uint_UintPtr(Uint64_Uint(Uint64Ptr_Uint64(i, opt...)))
}

func UintSlice_Uint64Slice(i []uint) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func UintSlice_Uint64SlicePtr(i []uint) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(UintSlice_Uint64Slice(i))
}
func Uint64Slice_UintSlice(i []uint64) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Uint64Slice_UintSlicePtr(i []uint64) *[]uint {
	return UintSlice_UintSlicePtr(Uint64Slice_UintSlice(i))
}
func UintSlicePtr_Uint64Slice(i *[]uint, opt ...Option) []uint64 {
	return UintSlice_Uint64Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Uint64SlicePtr(i *[]uint, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(UintSlice_Uint64Slice(UintSlicePtr_UintSlice(i, opt...)))
}
func Uint64SlicePtr_UintSlice(i *[]uint64, opt ...Option) []uint {
	return Uint64Slice_UintSlice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_UintSlicePtr(i *[]uint64, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Uint64Slice_UintSlice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Uint_Uint8(i uint) uint8 {
	return uint8(i)
}
func Uint_Uint8Ptr(i uint) *uint8 {
	return Uint8_Uint8Ptr(Uint_Uint8(i))
}
func Uint8_Uint(i uint8) uint {
	return uint(i)
}
func Uint8_UintPtr(i uint8) *uint {
	return Uint_UintPtr(Uint8_Uint(i))
}
func UintPtr_Uint8(i *uint, opt ...Option) uint8 {
	return Uint_Uint8(UintPtr_Uint(i, opt...))
}
func UintPtr_Uint8Ptr(i *uint, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Uint_Uint8(UintPtr_Uint(i, opt...)))
}
func Uint8Ptr_Uint(i *uint8, opt ...Option) uint {
	return Uint8_Uint(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_UintPtr(i *uint8, opt ...Option) *uint {
	return Uint_UintPtr(Uint8_Uint(Uint8Ptr_Uint8(i, opt...)))
}

func UintSlice_Uint8Slice(i []uint) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func UintSlice_Uint8SlicePtr(i []uint) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(UintSlice_Uint8Slice(i))
}
func Uint8Slice_UintSlice(i []uint8) []uint {
	res := []uint{}
	for _, item := range i {
		res = append(res, uint(item))
	}
	return res
}
func Uint8Slice_UintSlicePtr(i []uint8) *[]uint {
	return UintSlice_UintSlicePtr(Uint8Slice_UintSlice(i))
}
func UintSlicePtr_Uint8Slice(i *[]uint, opt ...Option) []uint8 {
	return UintSlice_Uint8Slice(UintSlicePtr_UintSlice(i, opt...))
}
func UintSlicePtr_Uint8SlicePtr(i *[]uint, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(UintSlice_Uint8Slice(UintSlicePtr_UintSlice(i, opt...)))
}
func Uint8SlicePtr_UintSlice(i *[]uint8, opt ...Option) []uint {
	return Uint8Slice_UintSlice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_UintSlicePtr(i *[]uint8, opt ...Option) *[]uint {
	return UintSlice_UintSlicePtr(Uint8Slice_UintSlice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Uint16_Uint32(i uint16) uint32 {
	return uint32(i)
}
func Uint16_Uint32Ptr(i uint16) *uint32 {
	return Uint32_Uint32Ptr(Uint16_Uint32(i))
}
func Uint32_Uint16(i uint32) uint16 {
	return uint16(i)
}
func Uint32_Uint16Ptr(i uint32) *uint16 {
	return Uint16_Uint16Ptr(Uint32_Uint16(i))
}
func Uint16Ptr_Uint32(i *uint16, opt ...Option) uint32 {
	return Uint16_Uint32(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Uint32Ptr(i *uint16, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Uint16_Uint32(Uint16Ptr_Uint16(i, opt...)))
}
func Uint32Ptr_Uint16(i *uint32, opt ...Option) uint16 {
	return Uint32_Uint16(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Uint16Ptr(i *uint32, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Uint32_Uint16(Uint32Ptr_Uint32(i, opt...)))
}

func Uint16Slice_Uint32Slice(i []uint16) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func Uint16Slice_Uint32SlicePtr(i []uint16) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Uint16Slice_Uint32Slice(i))
}
func Uint32Slice_Uint16Slice(i []uint32) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func Uint32Slice_Uint16SlicePtr(i []uint32) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Uint32Slice_Uint16Slice(i))
}
func Uint16SlicePtr_Uint32Slice(i *[]uint16, opt ...Option) []uint32 {
	return Uint16Slice_Uint32Slice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Uint32SlicePtr(i *[]uint16, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Uint16Slice_Uint32Slice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}
func Uint32SlicePtr_Uint16Slice(i *[]uint32, opt ...Option) []uint16 {
	return Uint32Slice_Uint16Slice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Uint16SlicePtr(i *[]uint32, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Uint32Slice_Uint16Slice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}

func Uint16_Uint64(i uint16) uint64 {
	return uint64(i)
}
func Uint16_Uint64Ptr(i uint16) *uint64 {
	return Uint64_Uint64Ptr(Uint16_Uint64(i))
}
func Uint64_Uint16(i uint64) uint16 {
	return uint16(i)
}
func Uint64_Uint16Ptr(i uint64) *uint16 {
	return Uint16_Uint16Ptr(Uint64_Uint16(i))
}
func Uint16Ptr_Uint64(i *uint16, opt ...Option) uint64 {
	return Uint16_Uint64(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Uint64Ptr(i *uint16, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Uint16_Uint64(Uint16Ptr_Uint16(i, opt...)))
}
func Uint64Ptr_Uint16(i *uint64, opt ...Option) uint16 {
	return Uint64_Uint16(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Uint16Ptr(i *uint64, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Uint64_Uint16(Uint64Ptr_Uint64(i, opt...)))
}

func Uint16Slice_Uint64Slice(i []uint16) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func Uint16Slice_Uint64SlicePtr(i []uint16) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Uint16Slice_Uint64Slice(i))
}
func Uint64Slice_Uint16Slice(i []uint64) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func Uint64Slice_Uint16SlicePtr(i []uint64) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Uint64Slice_Uint16Slice(i))
}
func Uint16SlicePtr_Uint64Slice(i *[]uint16, opt ...Option) []uint64 {
	return Uint16Slice_Uint64Slice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Uint64SlicePtr(i *[]uint16, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Uint16Slice_Uint64Slice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}
func Uint64SlicePtr_Uint16Slice(i *[]uint64, opt ...Option) []uint16 {
	return Uint64Slice_Uint16Slice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Uint16SlicePtr(i *[]uint64, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Uint64Slice_Uint16Slice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Uint16_Uint8(i uint16) uint8 {
	return uint8(i)
}
func Uint16_Uint8Ptr(i uint16) *uint8 {
	return Uint8_Uint8Ptr(Uint16_Uint8(i))
}
func Uint8_Uint16(i uint8) uint16 {
	return uint16(i)
}
func Uint8_Uint16Ptr(i uint8) *uint16 {
	return Uint16_Uint16Ptr(Uint8_Uint16(i))
}
func Uint16Ptr_Uint8(i *uint16, opt ...Option) uint8 {
	return Uint16_Uint8(Uint16Ptr_Uint16(i, opt...))
}
func Uint16Ptr_Uint8Ptr(i *uint16, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Uint16_Uint8(Uint16Ptr_Uint16(i, opt...)))
}
func Uint8Ptr_Uint16(i *uint8, opt ...Option) uint16 {
	return Uint8_Uint16(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Uint16Ptr(i *uint8, opt ...Option) *uint16 {
	return Uint16_Uint16Ptr(Uint8_Uint16(Uint8Ptr_Uint8(i, opt...)))
}

func Uint16Slice_Uint8Slice(i []uint16) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func Uint16Slice_Uint8SlicePtr(i []uint16) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Uint16Slice_Uint8Slice(i))
}
func Uint8Slice_Uint16Slice(i []uint8) []uint16 {
	res := []uint16{}
	for _, item := range i {
		res = append(res, uint16(item))
	}
	return res
}
func Uint8Slice_Uint16SlicePtr(i []uint8) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Uint8Slice_Uint16Slice(i))
}
func Uint16SlicePtr_Uint8Slice(i *[]uint16, opt ...Option) []uint8 {
	return Uint16Slice_Uint8Slice(Uint16SlicePtr_Uint16Slice(i, opt...))
}
func Uint16SlicePtr_Uint8SlicePtr(i *[]uint16, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Uint16Slice_Uint8Slice(Uint16SlicePtr_Uint16Slice(i, opt...)))
}
func Uint8SlicePtr_Uint16Slice(i *[]uint8, opt ...Option) []uint16 {
	return Uint8Slice_Uint16Slice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Uint16SlicePtr(i *[]uint8, opt ...Option) *[]uint16 {
	return Uint16Slice_Uint16SlicePtr(Uint8Slice_Uint16Slice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Uint32_Uint64(i uint32) uint64 {
	return uint64(i)
}
func Uint32_Uint64Ptr(i uint32) *uint64 {
	return Uint64_Uint64Ptr(Uint32_Uint64(i))
}
func Uint64_Uint32(i uint64) uint32 {
	return uint32(i)
}
func Uint64_Uint32Ptr(i uint64) *uint32 {
	return Uint32_Uint32Ptr(Uint64_Uint32(i))
}
func Uint32Ptr_Uint64(i *uint32, opt ...Option) uint64 {
	return Uint32_Uint64(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Uint64Ptr(i *uint32, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Uint32_Uint64(Uint32Ptr_Uint32(i, opt...)))
}
func Uint64Ptr_Uint32(i *uint64, opt ...Option) uint32 {
	return Uint64_Uint32(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Uint32Ptr(i *uint64, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Uint64_Uint32(Uint64Ptr_Uint64(i, opt...)))
}

func Uint32Slice_Uint64Slice(i []uint32) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func Uint32Slice_Uint64SlicePtr(i []uint32) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Uint32Slice_Uint64Slice(i))
}
func Uint64Slice_Uint32Slice(i []uint64) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func Uint64Slice_Uint32SlicePtr(i []uint64) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Uint64Slice_Uint32Slice(i))
}
func Uint32SlicePtr_Uint64Slice(i *[]uint32, opt ...Option) []uint64 {
	return Uint32Slice_Uint64Slice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Uint64SlicePtr(i *[]uint32, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Uint32Slice_Uint64Slice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}
func Uint64SlicePtr_Uint32Slice(i *[]uint64, opt ...Option) []uint32 {
	return Uint64Slice_Uint32Slice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Uint32SlicePtr(i *[]uint64, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Uint64Slice_Uint32Slice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}

func Uint32_Uint8(i uint32) uint8 {
	return uint8(i)
}
func Uint32_Uint8Ptr(i uint32) *uint8 {
	return Uint8_Uint8Ptr(Uint32_Uint8(i))
}
func Uint8_Uint32(i uint8) uint32 {
	return uint32(i)
}
func Uint8_Uint32Ptr(i uint8) *uint32 {
	return Uint32_Uint32Ptr(Uint8_Uint32(i))
}
func Uint32Ptr_Uint8(i *uint32, opt ...Option) uint8 {
	return Uint32_Uint8(Uint32Ptr_Uint32(i, opt...))
}
func Uint32Ptr_Uint8Ptr(i *uint32, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Uint32_Uint8(Uint32Ptr_Uint32(i, opt...)))
}
func Uint8Ptr_Uint32(i *uint8, opt ...Option) uint32 {
	return Uint8_Uint32(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Uint32Ptr(i *uint8, opt ...Option) *uint32 {
	return Uint32_Uint32Ptr(Uint8_Uint32(Uint8Ptr_Uint8(i, opt...)))
}

func Uint32Slice_Uint8Slice(i []uint32) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func Uint32Slice_Uint8SlicePtr(i []uint32) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Uint32Slice_Uint8Slice(i))
}
func Uint8Slice_Uint32Slice(i []uint8) []uint32 {
	res := []uint32{}
	for _, item := range i {
		res = append(res, uint32(item))
	}
	return res
}
func Uint8Slice_Uint32SlicePtr(i []uint8) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Uint8Slice_Uint32Slice(i))
}
func Uint32SlicePtr_Uint8Slice(i *[]uint32, opt ...Option) []uint8 {
	return Uint32Slice_Uint8Slice(Uint32SlicePtr_Uint32Slice(i, opt...))
}
func Uint32SlicePtr_Uint8SlicePtr(i *[]uint32, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Uint32Slice_Uint8Slice(Uint32SlicePtr_Uint32Slice(i, opt...)))
}
func Uint8SlicePtr_Uint32Slice(i *[]uint8, opt ...Option) []uint32 {
	return Uint8Slice_Uint32Slice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Uint32SlicePtr(i *[]uint8, opt ...Option) *[]uint32 {
	return Uint32Slice_Uint32SlicePtr(Uint8Slice_Uint32Slice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Uint64_Uint8(i uint64) uint8 {
	return uint8(i)
}
func Uint64_Uint8Ptr(i uint64) *uint8 {
	return Uint8_Uint8Ptr(Uint64_Uint8(i))
}
func Uint8_Uint64(i uint8) uint64 {
	return uint64(i)
}
func Uint8_Uint64Ptr(i uint8) *uint64 {
	return Uint64_Uint64Ptr(Uint8_Uint64(i))
}
func Uint64Ptr_Uint8(i *uint64, opt ...Option) uint8 {
	return Uint64_Uint8(Uint64Ptr_Uint64(i, opt...))
}
func Uint64Ptr_Uint8Ptr(i *uint64, opt ...Option) *uint8 {
	return Uint8_Uint8Ptr(Uint64_Uint8(Uint64Ptr_Uint64(i, opt...)))
}
func Uint8Ptr_Uint64(i *uint8, opt ...Option) uint64 {
	return Uint8_Uint64(Uint8Ptr_Uint8(i, opt...))
}
func Uint8Ptr_Uint64Ptr(i *uint8, opt ...Option) *uint64 {
	return Uint64_Uint64Ptr(Uint8_Uint64(Uint8Ptr_Uint8(i, opt...)))
}

func Uint64Slice_Uint8Slice(i []uint64) []uint8 {
	res := []uint8{}
	for _, item := range i {
		res = append(res, uint8(item))
	}
	return res
}
func Uint64Slice_Uint8SlicePtr(i []uint64) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Uint64Slice_Uint8Slice(i))
}
func Uint8Slice_Uint64Slice(i []uint8) []uint64 {
	res := []uint64{}
	for _, item := range i {
		res = append(res, uint64(item))
	}
	return res
}
func Uint8Slice_Uint64SlicePtr(i []uint8) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Uint8Slice_Uint64Slice(i))
}
func Uint64SlicePtr_Uint8Slice(i *[]uint64, opt ...Option) []uint8 {
	return Uint64Slice_Uint8Slice(Uint64SlicePtr_Uint64Slice(i, opt...))
}
func Uint64SlicePtr_Uint8SlicePtr(i *[]uint64, opt ...Option) *[]uint8 {
	return Uint8Slice_Uint8SlicePtr(Uint64Slice_Uint8Slice(Uint64SlicePtr_Uint64Slice(i, opt...)))
}
func Uint8SlicePtr_Uint64Slice(i *[]uint8, opt ...Option) []uint64 {
	return Uint8Slice_Uint64Slice(Uint8SlicePtr_Uint8Slice(i, opt...))
}
func Uint8SlicePtr_Uint64SlicePtr(i *[]uint8, opt ...Option) *[]uint64 {
	return Uint64Slice_Uint64SlicePtr(Uint8Slice_Uint64Slice(Uint8SlicePtr_Uint8Slice(i, opt...)))
}

func Float32_Float32Ptr(i float32) *float32 { return &i }
func Float32Ptr_Float32(i *float32, opt ...Option) float32 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(float32)
	}
	return *i
}

func Float32Slice_Float32SlicePtr(i []float32) *[]float32 { return &i }
func Float32SlicePtr_Float32Slice(i *[]float32, opt ...Option) []float32 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]float32)
	}
	return *i
}

func Float32_Float64(i float32) float64 {
	return float64(i)
}
func Float32_Float64Ptr(i float32) *float64 {
	return Float64_Float64Ptr(Float32_Float64(i))
}
func Float64_Float32(i float64) float32 {
	return float32(i)
}
func Float64_Float32Ptr(i float64) *float32 {
	return Float32_Float32Ptr(Float64_Float32(i))
}
func Float64_Float64Ptr(i float64) *float64 { return &i }
func Float32Ptr_Float64(i *float32, opt ...Option) float64 {
	return Float32_Float64(Float32Ptr_Float32(i, opt...))
}
func Float32Ptr_Float64Ptr(i *float32, opt ...Option) *float64 {
	return Float64_Float64Ptr(Float32_Float64(Float32Ptr_Float32(i, opt...)))
}
func Float64Ptr_Float32(i *float64, opt ...Option) float32 {
	return Float64_Float32(Float64Ptr_Float64(i, opt...))
}
func Float64Ptr_Float64(i *float64, opt ...Option) float64 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(float64)
	}
	return *i
}
func Float64Ptr_Float32Ptr(i *float64, opt ...Option) *float32 {
	return Float32_Float32Ptr(Float64_Float32(Float64Ptr_Float64(i, opt...)))
}

func Float32Slice_Float64Slice(i []float32) []float64 {
	res := []float64{}
	for _, item := range i {
		res = append(res, float64(item))
	}
	return res
}
func Float32Slice_Float64SlicePtr(i []float32) *[]float64 {
	return Float64Slice_Float64SlicePtr(Float32Slice_Float64Slice(i))
}
func Float64Slice_Float32Slice(i []float64) []float32 {
	res := []float32{}
	for _, item := range i {
		res = append(res, float32(item))
	}
	return res
}
func Float64Slice_Float32SlicePtr(i []float64) *[]float32 {
	return Float32Slice_Float32SlicePtr(Float64Slice_Float32Slice(i))
}
func Float64Slice_Float64SlicePtr(i []float64) *[]float64 { return &i }
func Float32SlicePtr_Float64Slice(i *[]float32, opt ...Option) []float64 {
	return Float32Slice_Float64Slice(Float32SlicePtr_Float32Slice(i, opt...))
}
func Float32SlicePtr_Float64SlicePtr(i *[]float32, opt ...Option) *[]float64 {
	return Float64Slice_Float64SlicePtr(Float32Slice_Float64Slice(Float32SlicePtr_Float32Slice(i, opt...)))
}
func Float64SlicePtr_Float32Slice(i *[]float64, opt ...Option) []float32 {
	return Float64Slice_Float32Slice(Float64SlicePtr_Float64Slice(i, opt...))
}
func Float64SlicePtr_Float64Slice(i *[]float64, opt ...Option) []float64 {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]float64)
	}
	return *i
}
func Float64SlicePtr_Float32SlicePtr(i *[]float64, opt ...Option) *[]float32 {
	return Float32Slice_Float32SlicePtr(Float64Slice_Float32Slice(Float64SlicePtr_Float64Slice(i, opt...)))
}

func String_StringPtr(i string) *string { return &i }
func StringPtr_String(i *string, opt ...Option) string {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(string)
	}
	return *i
}

func StringSlice_StringSlicePtr(i []string) *[]string { return &i }
func StringSlicePtr_StringSlice(i *[]string, opt ...Option) []string {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]string)
	}
	return *i
}

func Bool_BoolPtr(i bool) *bool { return &i }
func BoolPtr_Bool(i *bool, opt ...Option) bool {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new(bool)
	}
	return *i
}

func BoolSlice_BoolSlicePtr(i []bool) *[]bool { return &i }
func BoolSlicePtr_BoolSlice(i *[]bool, opt ...Option) []bool {
	if opt[0] == UseDefaultEmpty && i == nil {
		return *new([]bool)
	}
	return *i
}
