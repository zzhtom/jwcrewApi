package main

//for jwcrewApi
/**
*code list:
*	200 → success
*   204 → failed
*   205 → exist
*   404 → resource not found
*   422 → parameter error
 */
type jwcrewApi struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
