package day01

import println
import readInput
import kotlin.math.abs

fun main() {
    fun part1(input: List<String>): Int {
        val columns = getIntColumns(input)
        val firstColumn = columns.first.sorted()
        val secondColumn = columns.second.sorted()
        return firstColumn.zip(secondColumn) { first, second -> abs(first - second) }.sum()
    }

    fun part2(input: List<String>): Int {
        val columns = getIntColumns(input)
        val groupings = columns.second.groupingBy { it }.eachCount()
        return columns.first.map {
            groupings[it]?.times(it) ?: 0
        }.reduce(Int::plus)
    }

    check(part1(readInput("day1/part1_example")) == 11)
    check(part2(readInput("day1/part1_example")) == 31)

    val input = readInput("day1/part1")
    part1(input).println()
    part2(input).println()
}

fun getIntColumns(input: List<String>): Pair<List<Int>, List<Int>> {
    return input.map {
        val digits = it.split(" ")
        Pair(digits.first(), digits.last())
    }
        .map { (first, second) -> Pair(first.toInt(), second.toInt()) }
        .unzip()
}
