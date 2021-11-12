package com.uzabase

import io.kotest.core.spec.style.StringSpec
import io.mockk.mockk
import java.nio.file.Path

class FileWriterTest : StringSpec({
    // TODO
    "リクエストパスを書き込む" {
        val path = mockk<Path>()

        val sut = FileWriter(path)

        sut.writeRequestPath("/test")

        //TODO
//        /tmp/companyApi
//        - request.txt
//          POST / test?q = hoge
//          Content-Type: text/plain
//        - request_body.json
    }
})
