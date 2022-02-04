package com.uzabase.playcdc.internal.infra

import io.kotest.core.spec.style.StringSpec
import io.kotest.matchers.shouldBe

class HttpTest : StringSpec({
    "toOkHttpHeaders" {
        val result = toOkHttpHeaders(mapOf("key" to mapOf("valueKey" to "valueValue")))
        result.size shouldBe 1
        result.values("key").size shouldBe 1
        result.values("key")[0] shouldBe "valueValue"
    }
})
