package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.WireMock
import com.thoughtworks.gauge.BeforeScenario
import com.thoughtworks.gauge.BeforeSpec
import io.kotest.core.spec.style.StringSpec

class IntegrationTest : StringSpec({
    "integration" {
        callStoreMock()
    }
})

@BeforeScenario(tags = ["tagName"])
fun callStoreMock() {
    PlayCdc.storeMock(
        WireMock.get("/test")
            .withHeader("content-type", WireMock.equalTo("text/plain"))
            .withQueryParam("q", WireMock.equalTo("hey"))
            .withRequestBody(WireMock.equalTo("""{"key":"value"}"""))
            .willReturn(
                WireMock.aResponse()
                    .withHeader("content-type", "application/json")
                    .withBody("""{"count":1}""")))
}
