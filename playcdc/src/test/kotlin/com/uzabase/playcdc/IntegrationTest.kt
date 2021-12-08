package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.WireMock
import com.thoughtworks.gauge.BeforeScenario
import io.kotest.core.spec.style.StringSpec

class IntegrationTest : StringSpec({
    "call store mock" {
        callStoreMock()
    }

    "send request" {
        val json = """
            {
              "url" : "/test?q=hey",
              "method" : "GET",
              "headers" : {
                "content-type" : "text/plain"
              },
              "body" : {
                "key" : "value"
              }
            }
        """.trimIndent()
        PlayCdc.sendRequest(json)
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
                    .withStatus(200)
                    .withHeader("content-type", "application/json")
                    .withBody("""{"count":1}""")))
}
