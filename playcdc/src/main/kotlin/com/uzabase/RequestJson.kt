package com.uzabase

data class RequestJson(
    val url: String,
    val method: String,
    val header: Map<String, Any>,
    val body: String?
)
