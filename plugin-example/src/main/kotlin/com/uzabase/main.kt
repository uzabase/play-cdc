package com.uzabase

import io.grpc.netty.shaded.io.grpc.netty.NettyServerBuilder
import java.net.InetSocketAddress
import kotlin.system.exitProcess

fun main(args: Array<String>) {
    println("Hello World!")

    val server = NettyServerBuilder
        .forAddress(InetSocketAddress("localhost", 0))
        .executor(Runnable::run)
        .build()

    println("Starting server...")
    server.start()

    val port = server.port
    println("Listening on port: $port")

    server.awaitTermination();
    println("Terminated.")

    exitProcess(0)
}
