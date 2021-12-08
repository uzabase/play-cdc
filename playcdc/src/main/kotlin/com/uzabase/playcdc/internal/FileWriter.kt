package com.uzabase.playcdc.internal

import com.uzabase.playcdc.internal.infra.toPrettyString
import java.nio.file.Files
import java.nio.file.Path

class FileWriter(private val folderPath: Path) : Writer {
    override fun setup() {
        Files.createDirectories(folderPath)
    }

    override fun write(request: Request) {
        Files.write(folderPath.resolve("request.json"), toPrettyString(request).toByteArray())
    }

    override fun write(response: Response) {
        Files.write(folderPath.resolve("response.json"), toPrettyString(response).toByteArray())
    }
}
