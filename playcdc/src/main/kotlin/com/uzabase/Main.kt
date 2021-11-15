package com.uzabase

import com.github.tomakehurst.wiremock.client.WireMock
import com.thoughtworks.gauge.BeforeScenario

const val BASE_PATH = "/tmp"

fun main() {
    callStoreMock()
}

@BeforeScenario(tags = ["tagName"])
fun callStoreMock() {
    storeMock(WireMock.get("/test"))
}
