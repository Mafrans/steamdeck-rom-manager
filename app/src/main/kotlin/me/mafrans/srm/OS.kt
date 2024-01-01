package me.mafrans.srm

object OS {
    val name = System.getProperty("os.name")
    val isWindows = name.contains("Windows")
    val isUnix = name == "Unix" || name == "Linux"
    val isMac = name.contains("Mac")
}