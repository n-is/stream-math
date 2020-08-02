/* package sort defines various sorting algorithms for stream
 * data. It supports addition and removal of elements and
 * provides an iterator for accessing the elements in order.
 */
package sort

import "container/list"

type Sort interface {
	Add(v interface{})

	Remove(v interface{})

	First() *list.Element

	Last() *list.Element

	Reset()
}
