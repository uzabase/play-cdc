package com.uzabase.playcdc.internal.infra

import com.thoughtworks.gauge.BeforeScenario
import com.thoughtworks.gauge.BeforeSpec
import java.lang.reflect.Method

fun findFolderName(): String? = StackWalker.getInstance(StackWalker.Option.RETAIN_CLASS_REFERENCE)
    .walk { frames ->
        frames.map { f ->
            val methods = Class.forName(f.className).methods.filterNotNull().filter { it.name == f.methodName }
            methods.map(::toTags).flatten()
        }
            .filter { it.isNotEmpty() }
            .map { it.joinToString("_") }
            .findFirst().orElse(null)
    }

private fun toTags(method: Method): List<String> = method.annotations.map {
    when (it) {
        is BeforeScenario -> it.tags.toList()
        is BeforeSpec -> it.tags.toList()
        else -> emptyList()
    }
}.flatten()
