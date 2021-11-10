package com.uzabase

import com.thoughtworks.gauge.BeforeScenario
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe

class FunctionsTest : FreeSpec() {

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

            "複数タグの場合" {
                var folderName: String? = null
                MultipleTagNames().callStoreMock { f -> folderName = f }
                folderName shouldBe "tagName_otherTagName"
            }

            "同名で別シグニチャのメソッドがある場合" {
                var folderName: String? = null
                SameName().callStoreMock { f -> folderName = f }
                folderName shouldBe "tagName"
            }

            "該当するアノテーションがない場合" {
                NoAnnotation().callStoreMock { throw RuntimeException("should not be called") }
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

    class MultipleTagNames {
        @BeforeScenario(tags = ["tagName", "otherTagName"])
        fun callStoreMock(writer: (String) -> Unit) {
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

    class NoAnnotation {
        fun callStoreMock(writer: (String) -> Unit) {
            storeMock(writer)
        }
    }
}
