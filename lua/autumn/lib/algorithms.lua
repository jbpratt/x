function stack()
  local out = {}
  out.push = function(item)
    out[#out+1] = item
  end
  out.pop = function()
    if #out>0 then
      return table.remove(out, #out)
    end
  end
  out.interator = function()
    return function()
      return out.pop()
    end
  end
  return out
end

function queue()
  local out = {}
  local first, last = 0, -1
  out.push = function(item)
    last = last + 1
    out[last] = item
  end
  out.pop = function()
    if first <= last then
      local val = out[first]
      out[first] = nil
      first = first + 1
      return val
    end
  end
  out.interator = function()
    return function()
      return out.pop()
    end
  end
  setmetatable(out, {
      _len = function()
        return (last-first+1)
      end,
    })
  return out
end

return {
  stack = stack,
  queue = queue,
}
