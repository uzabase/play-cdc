package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.WireMockServer
import com.github.tomakehurst.wiremock.client.WireMock.*
import com.github.tomakehurst.wiremock.core.WireMockConfiguration.wireMockConfig
import io.kotest.core.spec.style.StringSpec

class IntegrationTest : StringSpec({

    lateinit var wiremock: WireMockServer

    beforeSpec {
        wiremock = WireMockServer(wireMockConfig().port(8080))
        wiremock.start()
    }

    afterEach {
        wiremock.resetAll()
    }

    afterSpec {
        wiremock.stop()
    }

    "send request" {
        val contract = """
            {
              "request": {
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

        PlayCdc.sendRequest("http://localhost:8080", contract)

        wiremock.verify(getRequestedFor(urlPathEqualTo("/test"))
            .withQueryParam("q", equalTo("hey"))
            .withHeader("content-type", equalTo("application/json")))
    }

    "send request by `urlPath`" {
        val contract = """
            {
              "request": {
                "urlPath": "/test",
                "method": "GET"
              },
              "response": {
                "status": 200
              }
            }
        """.trimIndent()
        PlayCdc.sendRequest("http://localhost:8080", contract)

        wiremock.verify(getRequestedFor(urlPathEqualTo("/test")))
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

        wiremock.verify(putRequestedFor(urlPathEqualTo("/test"))
            .withQueryParam("q", equalTo("hey"))
            .withHeader("content-type", equalTo("application/json"))
            .withRequestBody(equalToJson(body)))
    }

    "verify response" {
        val json = """
            {
              "request": {
                "url": "/test?q=hey",
                "method": "GET",
                "headers": {
                  "content-type": {
                    "equalTo": "text/plain"
                  }
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
