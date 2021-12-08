package com.uzabase.playcdc.internal

import com.github.tomakehurst.wiremock.client.WireMock.*
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe
import io.mockk.clearAllMocks

class MappingBuilderExtensionsTest : FreeSpec({
    afterTest {
        clearAllMocks()
    }

    "MappingBuilderをRequestJsonに変換する" - {
        "URL" {
            val sut = get("/test")
            sut.toRequest().url shouldBe "/test"
        }

        "METHOD" - {
            "GET" {
                val sut = get("/test")
                sut.toRequest().method shouldBe "GET"
            }
            "POST" {
                val sut = post("/test")
                sut.toRequest().method shouldBe "POST"
            }
        }

        "Query Params" {
            val sut = get("/test")
                .withQueryParams(mapOf(
                    "p1" to equalTo("v1"),
                    "p2" to equalTo("v2")
                ))
            sut.toRequest().url shouldBe "/test?p1=v1&p2=v2"
        }

        "Headers" {
            val sut = get("/test")
                .withHeader("content-type", equalTo("text/plain"))
                .withHeader("Accept", equalTo("*/*"))

            sut.toRequest().headers shouldBe mapOf("content-type" to "text/plain", "Accept" to "*/*")
        }

        "Body" - {
            "With body" {
                val sut = post("/test")
                    .withRequestBody(equalToJson("""{
                            "total_results": 4,
                            "results": {
                                "key": "value"
                            }
                        }""".trimIndent()))

                sut.toRequest().body shouldBe mapOf(
                    "total_results" to 4,
                    "results" to mapOf(
                        "key" to "value"
                    ))
            }

            "With no body" {
                val sut = get("/test")

                sut.toRequest().body shouldBe emptyMap()
            }
        }
    }

    "MappingBuilderをResponseJsonに変換する" - {
        "Status code" {
            val sut = get("/test")
                .willReturn(aResponse().withStatus(200))

            sut.toResponse().status shouldBe 200
        }

        "Headers" {
            val sut = get("/test")
                .willReturn(aResponse()
                    .withHeader("content-type", "application/json"))

            sut.toResponse().headers shouldBe mapOf("content-type" to "application/json")
        }

        "Body" - {
            "With body" {
                val sut = post("/test")
                    .willReturn(aResponse()
                        .withBody("""{
                                "total_results": 4,
                                "results": {
                                    "key": "value"
                                }
                            }""".trimIndent()))

                sut.toResponse().body shouldBe mapOf(
                    "total_results" to 4,
                    "results" to mapOf(
                        "key" to "value"
                    ))
            }

            "With no body" {
                val sut = get("/test")

                sut.toResponse().body shouldBe emptyMap()
            }
        }
    }
})
