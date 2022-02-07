package com.uzabase.playcdc.internal

data class Contract(
    val request: Request,
    val response: Response
) {
    data class Request(
        val url: String?,
        val urlPath: String?,
        val method: String,
        val headers: Map<String, Map<String, String>>?
    ) {
        init {
            if (url == null && urlPath == null) {
                throw RuntimeException("Either `url` or `urlPath` is required")
            }
        }
    }

    data class Response(
        val status: Int,
        val headers: Map<String, String>?,
        val jsonBody: Map<String, Any>?
    )
}
