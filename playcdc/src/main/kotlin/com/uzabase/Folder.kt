package com.uzabase

import com.thoughtworks.gauge.BeforeScenario

fun getFolderName(): String? {
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
