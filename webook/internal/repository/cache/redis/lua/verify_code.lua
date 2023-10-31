local key = KEY[1]
local code = ARGV[1]
local cntKey = key..':cnt'
local targetVal = redis.call('get', key)
local cnt = tonumber(redis.call('get', cntKey))

if cnt <= 0 then
    return -1
elseif targetVal == code then
    redis.call('del', key)
    redis.call('del', cntKey)
    return 0
else
    redis.call('decr', cntKey, -1)
    return -2
end
