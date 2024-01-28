#include <string>
#include "strconv2.h"

extern "C"
__declspec(dllexport)
const wchar_t *greeting(const wchar_t *name) {
  static thread_local std::wstring result;
  result = format(L"Hello ハロー©: %s", name);
  return result.c_str();
}

extern "C"
__declspec(dllexport)
const char *greeting8(const char *name) {
  static thread_local std::string result;
  result = format(u8"Hello ハロー©: %s", name);
  return result.c_str();
}
