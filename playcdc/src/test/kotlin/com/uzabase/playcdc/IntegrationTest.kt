package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.WireMockServer
import com.github.tomakehurst.wiremock.client.WireMock
import com.github.tomakehurst.wiremock.core.WireMockConfiguration.wireMockConfig
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

    lateinit var wiremock: WireMockServer

    beforeSpec {
        wiremock = WireMockServer(wireMockConfig().port(8080))
        wiremock.start()
    }

    afterSpec {
        wiremock.stop()
    }

    "call store mock" {
        callStoreMock()
    }

    "store mock" {
        PlayCdc.storeMock(mappingBuilder, "company_api")
    }

    "send request" {
        val contract = """
            {
              "request": {
                "url": "/test?q=hey",
                "method": "GET",
                "headers": {
                  "content-type": {
                    "equalTo": "application/json"
                  }
                }
              },
              "response": {
                "status": 200
              }
            }
        """.trimIndent()

        val (status, body, headers) = PlayCdc.sendRequest("http://localhost:8080", contract)
        println(status)
        println(body)
        println(headers)
    }

    "send request with body" {
        val contract = """
            {
              "request": {
                "url": "/test?q=hey",
                "method": "PUT",
                "headers": {
                  "content-type": {
                    "equalTo": "application/json"
                  }
                }
              },
              "response": {
                "status": 200
              }
            }
        """.trimIndent()

        val body = """
            {
              "key": "value"
            }
        """.trimIndent()

        PlayCdc.sendRequest("http://localhost:8080", contract, body)
    }

    "verify response" {
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
                "status": 200,
                "headers": {
                  "content-type": "application/json"
                },
                "jsonBody": {
                    "key": "value"
                }
              }
            }
        """.trimIndent()

        val actualBody = """
            {
                "key": "value"
            }
        """.trimIndent()

        val actualHeaders = mapOf(
            "content-type" to "application/json",
            "content-length" to "18"
        )

        PlayCdc.verifyResponse(json, 200, actualBody, actualHeaders)
    }
})

@BeforeScenario(tags = ["tagName"])
fun callStoreMock() {
    PlayCdc.storeMock(mappingBuilder)
}
