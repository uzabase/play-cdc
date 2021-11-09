package com.uzabase

import com.thoughtworks.gauge.BeforeScenario
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe

class MainTest : FreeSpec() {

    init {
        "タグ名のついたフォルダを生成する" - {
            "コールスタックが単一階層の場合" {
                var folderName: String? = null
                SingleStack().callStoreMock { f -> folderName = f }
                folderName shouldBe "tagName"
            }

            "コールスタックが複数階層に渡る場合" {
                var folderName: String? = null
                MultipleStacks().callCallStoreMock { f -> folderName = f }
                folderName shouldBe "tagName"
            }

            "同名で別シグニチャのメソッドがある場合" {
                var folderName: String? = null
                SameName().callStoreMock { f -> folderName = f }
                folderName shouldBe "tagName"
            }
        }
    }

    class SingleStack {
        @BeforeScenario(tags = ["tagName"])
        fun callStoreMock(writer: (String) -> Unit) {
            storeMock(writer)
        }
    }

    class MultipleStacks {
        @BeforeScenario(tags = ["tagName"])
        fun callCallStoreMock(writer: (String) -> Unit) {
            callStoreMock(writer)
        }

        private fun callStoreMock(writer: (String) -> Unit) {
            storeMock(writer)
        }
    }

    class SameName {
        fun callStoreMock() {
            // do nothing
        }

        @BeforeScenario(tags = ["tagName"])
        fun callStoreMock(writer: (String) -> Unit) {
            storeMock(writer)
        }
    }
}
