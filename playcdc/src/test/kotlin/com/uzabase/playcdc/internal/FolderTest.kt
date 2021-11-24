package com.uzabase.playcdc.internal

import com.thoughtworks.gauge.BeforeScenario
import com.thoughtworks.gauge.BeforeSpec
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe

class FolderTest : FreeSpec({
    "タグ名からフォルダを取得する" - {
        "コールスタックが単一階層の場合" {
            SingleStack().callFindFolderName() shouldBe "tagName"
        }

        "コールスタックが複数階層に渡る場合" {
            MultipleStacks().callCallFindFolderName() shouldBe "tagName"
        }

        "複数タグの場合" {
            MultipleTagNames().callFindFolderName() shouldBe "tagName_otherTagName"
        }

        "同名で別シグニチャのメソッドがある場合" {
            SameName().callFindFolderName() shouldBe "tagName"
        }

        "該当するアノテーションがない場合" {
            NoAnnotation().callFindFolderName() shouldBe null
        }

        "BeforeScenarioだけでなく、BeforeSpecでも取得できる" {
            BeforeSpecCase().callFindFolderName() shouldBe "tagName"
        }
    }
})

class SingleStack {
    @BeforeScenario(tags = ["tagName"])
    fun callFindFolderName() = findFolderName()
}

class MultipleStacks {
    @BeforeScenario(tags = ["tagName"])
    fun callCallFindFolderName() = callFindFolderName()

    private fun callFindFolderName() = findFolderName()
}

class MultipleTagNames {
    @BeforeScenario(tags = ["tagName", "otherTagName"])
    fun callFindFolderName() = findFolderName()
}

class SameName {
    fun callFindFolderName(@Suppress("UNUSED_PARAMETER") dummy: String) = findFolderName()

    @BeforeScenario(tags = ["tagName"])
    fun callFindFolderName() = findFolderName()
}

class NoAnnotation {
    fun callFindFolderName() = findFolderName()
}

class BeforeSpecCase {
    @BeforeSpec(tags = ["tagName"])
    fun callFindFolderName() = findFolderName()
}
