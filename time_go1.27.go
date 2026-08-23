//go:build go1.27

package timekit

func (t Time[T]) As[U Timezone]() Time[U] {
	return Time[U]{utc: t.utc}
}
