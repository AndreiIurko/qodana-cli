package org.jetbrains.qodana

object QodanaPluginConstants {
    const val GROUP_NAME = "qodana"
    const val EXTENSION_NAME = "qodana"
    const val EXECUTABLE = "docker"

    const val RUN_INSPECTIONS_TASK_NAME = "runInspections"
    const val STOP_INSPECTIONS_TASK_NAME = "stopInspections"
    const val CLEAN_INSPECTIONS_TASK_NAME = "cleanInspections"

    const val DOCKER_CONTAINER_NAME_INSPECTIONS = "idea-inspections"
    const val DOCKER_IMAGE_NAME_INSPECTIONS = "jetbrains/qodana:latest"
}

