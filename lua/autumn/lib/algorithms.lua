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

function pqueue()
  local t = {}

  -- set of elements
  local set = {}
  local r_set = {}
  local keys = {}

  local function addKV(k, v)
    set[k] = v
    if not r_set[v] then
      ti(keys,v)
      table.sort(keys)
      local k0 = {k}
      r_set[v] = k0
      setmetatable(k0,{
          __mode = 'v'
        })
    else
      ti(r_set[v],k)
    end
  end

  local remove = function(k)
    local v = set[k]
    local prioritySet = r_set[v]
    tr2(prioritySet,k)
    if #prioritySet < 1 then
      tr2(keys,v)
      r_set[v] = nil
      table.sort(keys)
      set[k] = nil
    end
  end; t.remove = remove

  t.min = function()
    local priority = keys[1]
    if priority then
      return r_set[priority][1] or {}
    else
      return {}
    end
  end
  
  t.max = function()
    local priority = keys[#keys]
    if priority then
      return r_set[priority][1] or {}
    else
      return {}
    end
  end

  t.empty = function()
    return #keys < 1
  end

  t.iterate = function()
    return function()
      if not t.empty() then
        local ele = t.max()
        t.remove(ele)
        return ele
      end
    end
  end

  setmetatable(t, {
    __index = set,
    __newindex = function(t,k,v)
      if not set[k] then
        addKV(k,v)
      else
        remove(k)
        addKV(k,v)
      end
    end,
  })
  return t
end

return {
  stack = stack,
  queue = queue,
  pqueue = pqueue,
}
