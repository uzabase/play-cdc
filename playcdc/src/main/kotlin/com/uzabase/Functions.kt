package com.uzabase

import com.thoughtworks.gauge.BeforeScenario
import kotlin.io.path.Path
import kotlin.io.path.createDirectory

fun storeMock() {
    storeMock { f -> Path(BASE_PATH).resolve(f).createDirectory() }
}

internal fun storeMock(writer: (String) -> Unit) {
    val instance = StackWalker.getInstance(StackWalker.Option.RETAIN_CLASS_REFERENCE)
    instance.walk { frames ->
        frames.map { f ->
            val methods = Class.forName(f.className).methods.filterNotNull().filter { it.name == f.methodName }
            methods.mapNotNull { it.getAnnotation(BeforeScenario::class.java) }.map { it.tags.toList() }.flatten()
        }
            .filter { it.isNotEmpty() }
            .forEach { writer(it.joinToString("_")) }
    }
}
