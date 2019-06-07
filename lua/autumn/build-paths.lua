require "os_name"

local paths = {
  '{root}/{module}',
  '{root}/lib/{module}',
  '{root}/lib/external/{module}',
  '{root}/lib/external/platform-specific/{platform}/{module}',
}

local module_paths = {
  '?.{extension}',
  '?./init.{extension}',
  '?/core.{extension}',
}

local extensions = {
  Windows = 'dll',
  Linux = 'so',
  Mac = 'dylib',
}


local root_dir = '.'
local curr_platform, curr_arch = getOS()

local cpaths, lpaths = {}, {}
local curr_clib_ext = extensions[curr_platform]

if curr_clib_ext then
  for _, path in ipairs(paths) do
    local path = path:gsub("{(%w+)}", {
      root = root_dir,
      platform = curr_platform,
    })
    print(path)
  end
end

