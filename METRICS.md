# snap collector plugin - netstat

## Collected Metrics
http://git.kernel.org/cgit/linux/kernel/git/torvalds/linux.git/tree/include/net/tcp_states.h?id=HEAD

This plugin has the ability to gather the following metrics:

Namespace                               | Description
----------------------------------------|-------------
/staples/procfs/netstat/tcp_close       | Number of TCP Connection with no state (Should be Zero)
/staples/procfs/netstat/tcp_close_wait  | Connections waiting for a termination request from the local user
/staples/procfs/netstat/tcp_closing     | Connections waiting for a termination request from acknowledgment from the remote TCP
/staples/procfs/netstat/tcp_established | Established tcp connections
/staples/procfs/netstat/tcp_fin_wait1   | Connections waiting for a connection termination request from a remote TCP, or an acknowledgment of the connection termination request previously sent.
/staples/procfs/netstat/tcp_fin_wait2   | Connections waiting for a connection termination request from the remote TCP
/staples/procfs/netstat/tcp_last_ack    | Connections waiting for an acknowledgment of the connection termination previously sent to the remote TCP
/staples/procfs/netstat/tcp_listen      | Connections waiting for a connection request from any remote TCP and port
/staples/procfs/netstat/tcp_none        | Connections with inactive TCP
/staples/procfs/netstat/tcp_syn_recv    | Connections waiting for confirming connection request acknowledgment after having both received and sent a connection request
/staples/procfs/netstat/tcp_syn_sent    | Connections waiting for a matching connection request after having sent a connection request
/staples/procfs/netstat/tcp_time_wait   | Connections waiting to ensure the remote TCP received the acknowledgment of its connection termination request
/staples/procfs/netstat/udp_socket      | Number of UDP connections
