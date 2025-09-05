cmake_minimum_required(VERSION 3.15)
project(hexlib VERSION 0.1.0 LANGUAGES C CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_library(hexlib SHARED
    src/hexlib.cpp
    src/mpc_kalman.cpp
)

target_include_directories(hexlib
    PUBLIC
        ${CMAKE_CURRENT_SOURCE_DIR}/include
)

if(WIN32)
    target_compile_definitions(hexlib PRIVATE HEXLIB_EXPORTS)
endif()

set_target_properties(hexlib PROPERTIES
    OUTPUT_NAME "hexlib"
    PUBLIC_HEADER "include/hexlib/hexlib.h"
)

install(TARGETS hexlib
    RUNTIME DESTINATION bin
    LIBRARY DESTINATION lib
    ARCHIVE DESTINATION lib
    PUBLIC_HEADER DESTINATION include/hexlib
)


