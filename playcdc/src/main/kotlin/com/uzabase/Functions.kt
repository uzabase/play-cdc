package com.uzabase

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.thoughtworks.gauge.BeforeScenario
import kotlin.io.path.Path
import kotlin.io.path.createDirectory

interface Writer {
    fun writePath(path: String)
}

fun storeMock() {
    storeMock(TODO(), TODO())
}

internal fun storeMock(mappingBuilder: MappingBuilder, writer: Writer) {
    getFolderName()?.let {
        Path(BASE_PATH).resolve(it).createDirectory()
    }
}

internal fun getFolderName(): String? {
    return StackWalker.getInstance(StackWalker.Option.RETAIN_CLASS_REFERENCE).walk { frames ->
        frames.map { f ->
            val methods = Class.forName(f.className).methods.filterNotNull().filter { it.name == f.methodName }
            methods.mapNotNull { it.getAnnotation(BeforeScenario::class.java) }.map { it.tags.toList() }.flatten()
        }
            .filter { it.isNotEmpty() }
            .map { it.joinToString("_") }
            .findFirst().orElse(null)
    }
}
