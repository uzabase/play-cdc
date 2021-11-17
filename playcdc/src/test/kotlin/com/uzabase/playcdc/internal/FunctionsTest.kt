package com.uzabase.playcdc.internal

import com.github.tomakehurst.wiremock.client.WireMock.*
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe
import io.mockk.clearAllMocks

class FunctionsTest : FreeSpec({
    afterTest {
        clearAllMocks()
    }

    "MappingBuilderをRequestJsonに変換する" - {
        "URL" {
            val mappingBuilder = get("/test")
            toRequestJson(mappingBuilder).url shouldBe "/test"
        }

        "METHOD" - {
            "GET" {
                val mappingBuilder = get("/test")
                toRequestJson(mappingBuilder).method shouldBe "GET"
            }
            "POST" {
                val mappingBuilder = post("/test")
                toRequestJson(mappingBuilder).method shouldBe "POST"
            }
        }

        "Query Params" {
            val mappingBuilder = get("/test")
                .withQueryParams(mapOf(
                    "p1" to equalTo("v1"),
                    "p2" to equalTo("v2")
                ))
            toRequestJson(mappingBuilder).url shouldBe "/test?p1=v1&p2=v2"
        }

        "Headers" {
            val mappingBuilder = get("/test")
                .withHeader("content-type", equalTo("text/plain"))
                .withHeader("Accept", equalTo("*/*"))

            toRequestJson(mappingBuilder).headers shouldBe mapOf("content-type" to "text/plain", "Accept" to "*/*")
        }

        "Body" - {
            "With body" {
                val mappingBuilder = post("/test")
                    .withRequestBody(equalToJson("""{
                            "total_results": 4,
                            "results": {
                                "key": "value"
                            }
                        }""".trimIndent()))

                toRequestJson(mappingBuilder).body shouldBe mapOf(
                    "total_results" to 4,
                    "results" to mapOf(
                        "key" to "value"
                    ))
            }

            "With no body" {
                val mappingBuilder = get("/test")

                toRequestJson(mappingBuilder).body shouldBe emptyMap()
            }
        }
    }

    "MappingBuilderをResponseJsonに変換する" - {
        "Headers" {
            val mappingBuilder = get("/test")
                .willReturn(aResponse()
                    .withHeader("content-type", "application/json"))

            toResponseJson(mappingBuilder).headers shouldBe mapOf("content-type" to "application/json")
        }

        "Body" - {
            "With body" {
                val mappingBuilder = post("/test")
                    .willReturn(aResponse()
                        .withBody("""{
                                "total_results": 4,
                                "results": {
                                    "key": "value"
                                }
                            }""".trimIndent()))

                toResponseJson(mappingBuilder).body shouldBe mapOf(
                    "total_results" to 4,
                    "results" to mapOf(
                        "key" to "value"
                    ))
            }

            "With no body" {
                val mappingBuilder = get("/test")

                toResponseJson(mappingBuilder).body shouldBe emptyMap()
            }
        }
    }
})
