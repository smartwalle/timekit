//go:build go1.27

package timekit

func (t Time[T]) As[U Timezone](_ U) Time[U] {
	return Time[U]{utc: t.utc}
}
