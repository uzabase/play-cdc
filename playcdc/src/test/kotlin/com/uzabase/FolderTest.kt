import com.thoughtworks.gauge.BeforeScenario
import com.uzabase.*
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe

class FolderTest : FreeSpec({
    "タグ名のついたフォルダを生成する" - {
        "コールスタックが単一階層の場合" {
            SingleStack().callGetFolderName() shouldBe "tagName"
        }

        "コールスタックが複数階層に渡る場合" {
            MultipleStacks().callCallGetFolderName() shouldBe "tagName"
        }

        "複数タグの場合" {
            MultipleTagNames().callGetFolderName() shouldBe "tagName_otherTagName"
        }

        "同名で別シグニチャのメソッドがある場合" {
            SameName().callGetFolderName() shouldBe "tagName"
        }

        "該当するアノテーションがない場合" {
            NoAnnotation().callGetFolderName() shouldBe null
        }
    }
})

class SingleStack {
    @BeforeScenario(tags = ["tagName"])
    fun callGetFolderName() = getFolderName()
}

class MultipleStacks {
    @BeforeScenario(tags = ["tagName"])
    fun callCallGetFolderName() = callGetFolderName()

    private fun callGetFolderName() = getFolderName()
}

class MultipleTagNames {
    @BeforeScenario(tags = ["tagName", "otherTagName"])
    fun callGetFolderName() = getFolderName()
}

class SameName {
    fun callGetFolderName(dummy: String) = getFolderName()

    @BeforeScenario(tags = ["tagName"])
    fun callGetFolderName() = getFolderName()
}

class NoAnnotation {
    fun callGetFolderName() = getFolderName()
}
