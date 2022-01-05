package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.WireMock
import com.thoughtworks.gauge.BeforeScenario
import io.kotest.core.spec.style.StringSpec

private val mappingBuilder = WireMock.get("/test")
    .withHeader("content-type", WireMock.equalTo("text/plain"))
    .withQueryParam("q", WireMock.equalTo("hey"))
    .withRequestBody(WireMock.equalTo("""{"key":"value"}"""))
    .willReturn(
        WireMock.aResponse()
            .withStatus(200)
            .withHeader("content-type", "application/json")
            .withBody("""{"count":1}""")
    )

class IntegrationTest : StringSpec({
    "call store mock" {
        callStoreMock()
    }

    "store mock" {
        PlayCdc.storeMock(mappingBuilder, "company_api")
    }

    "send request" {
        val json = """
            {
              "request": {
                "url": "/test?q=hey",
                "method": "GET",
                "headers": {
                  "content-type": "text/plain"
                }
              },
              "response": {
                "status": 200
              }
            }
        """.trimIndent()

        PlayCdc.sendRequest("http://localhost:8080", json)
    }

    "verify response" {
        val json = """
            {
                "status": 200,
                "headers": {},
                "body": {
                    "key": "value"
                }
            }
        """.trimIndent()

        val expectedBody = """
            {
                "key": "value"
            }
        """.trimIndent()

        PlayCdc.verifyResponse(json, 200, expectedBody)
    }
})

@BeforeScenario(tags = ["tagName"])
fun callStoreMock() {
    PlayCdc.storeMock(mappingBuilder)
}