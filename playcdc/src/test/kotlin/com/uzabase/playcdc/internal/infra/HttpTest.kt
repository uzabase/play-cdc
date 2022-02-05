package com.uzabase.playcdc.internal.infra

import io.kotest.core.spec.style.StringSpec
import io.kotest.matchers.shouldBe

class HttpTest : StringSpec({
    "toArray" {
        toArray(mapOf("key" to mapOf("valueKey" to "valueValue"))) shouldBe arrayOf("key", "valueValue")
    }
})
