package com.uzabase

import com.thoughtworks.gauge.BeforeScenario

const val BASE_PATH = "/tmp"

fun main() {
    callStoreMock()
}

@BeforeScenario(tags = ["tagName"])
fun callStoreMock() {
    storeMock()
}
