package com.uzabase.playcdc.internal.legacy

data class Request(
    val url: String,
    val method: String,
    val headers: Map<String, String>,
    val body: Map<String, Any>?
) {
    val contentType: String? = headers["content-type"]
}
