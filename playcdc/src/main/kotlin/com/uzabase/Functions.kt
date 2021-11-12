package com.uzabase

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.thoughtworks.gauge.BeforeScenario
import kotlin.io.path.Path
import kotlin.io.path.createDirectory

interface Writer {
    fun writePath(path: String)
}

fun forTest(): String? {
    return getFolderName()
}

fun storeMock() {
    storeMock(TODO(), TODO())
}

internal fun storeMock(mappingBuilder: MappingBuilder, writer: Writer) {
    getFolderName()?.let {
        Path(BASE_PATH).resolve(it).createDirectory()
    }
}
