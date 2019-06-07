function getOS()
  os_name, arch_name = '', ''

	if jit and jit.os and jit.arch then
		os_name = jit.os
    arch_name = jit.arch
  else
    local popen_status, popen_result = pcall(io.popen, "")
		if popen_status then
			popen_result:close()
			-- Unix-based OS
			os_name = io.popen('uname -s','r'):read('*l')
      arch_name = io.popen('uname -m','r'):read('*l')
    else
      print('not supported')
    end
  end

  return os_name, arch_name
end
