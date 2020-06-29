package _map


type Map interface {

	/**
		获取map大小
	 */
	size() (val int, err error)

	/**
		判断map类型是否为空
	 */
	isEmpty() (val bool, err error)

	/**
		判断是否含有key
	 */
	containsKey() (val bool, err error)

	/**
		判断是否含有value
	 */
	containsValue() (val bool, err error)
}