# 目标  

> 意图可以从 1 ~ N 台服务器中选举出的 leader   

# 调查  
> 发现 raft, bully 算法或调用 k8s 的 client-go 算法都需要调用大量库  
> 感觉用法太重，而需求比较简单，因此弃用上述算法      

# 意图  
> 如果使用 mysql 则可以通过 select for update 创建存储过程，直接利用 mysql 自身锁  
> 因项目不允许使用 mysql 存储过程，因此放弃 mysql 作为选举媒介   
> 通过 redis 对某个 key 进行锁抢占方法，获得锁的服务器则为 leader  
> 通过 reflesh TTL 的方法实现锁延时, 从而锁定 leader   

# 限制  
> 服务器选择 leader 以抢占 redis 锁为目标   
> 没有必须大于 2 台服务器才可以选择 leader 限制  


# 用法 example  
> 1. 先抢占 key 锁，则设定 ttl 为 5s, 每秒更新一次锁 ttl, 节点为 leader       
> 2. 其他服务器访问 key 时，发现无法获取 Key 锁，则 3s 后重试, 节点为 slave    
