package com.uzabase

import com.github.tomakehurst.wiremock.client.WireMock
import com.github.tomakehurst.wiremock.client.WireMock.*
import com.github.tomakehurst.wiremock.matching.StringValuePattern
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.equality.beEqualToUsingFields
import io.kotest.matchers.shouldBe
import io.mockk.clearAllMocks

class FunctionsTest : FreeSpec({
    afterTest {
        clearAllMocks()
    }

    "MappingBuilderをRequestJsonに変換する" - {
        "URL" {
            val mappingBuilder = WireMock.get("/test")
            toRequestJson(mappingBuilder).url shouldBe "/test"
        }

        "METHOD" - {
            "GET" {
                val mappingBuilder = WireMock.get("/test")
                toRequestJson(mappingBuilder).method shouldBe "GET"
            }
            "POST" {
                val mappingBuilder = WireMock.post("/test")
                toRequestJson(mappingBuilder).method shouldBe "POST"
            }
        }

        "Query Params" {
            val mappingBuilder = WireMock.get("/test")
                .withQueryParams(mapOf(
                    "p1" to equalTo("v1"),
                    "p2" to equalTo("v2")
                ))
            toRequestJson(mappingBuilder).url shouldBe "/test?p1=v1&p2=v2"
        }

        "Header" {
            val mappingBuilder = WireMock.get("/test")
                .withHeader("content-type", equalTo("text/plain"))
                .withHeader("Accept", equalTo("*/*"))

            toRequestJson(mappingBuilder).header shouldBe mapOf("content-type" to "text/plain", "Accept" to "*/*")
        }

        "Body" - {
            "With body" {
                val mappingBuilder = WireMock.post("/test")
                    .withRequestBody(equalToJson("{ \"total_results\": 4 }"))

                toRequestJson(mappingBuilder).body shouldBe "{ \"total_results\": 4 }"
            }

            "With no body" {
                val mappingBuilder = WireMock.get("/test")

                toRequestJson(mappingBuilder).body shouldBe null
            }
        }
    }
})
