package com.uzabase.playcdc.internal

data class Response(
    val status: Int,
    val headers: Map<String, String>,
    val body: Map<String, Any>
)
