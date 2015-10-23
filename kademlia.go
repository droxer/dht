package kademlia

const (
	// alpha is a small number representing the degree of parallelism in network calls, usually
	alpha = 3
	// B is the size in bits of the keys used to identify nodes and store and retrieve data;
	// in basic Kademlia this is 160, the length of an SHA1 digest (hash)
	B = 160
	// k is the maximum number of contacts stored in a bucket; this is normally 20
	K = 20
)

const (
	// expire is the time after which a key/value pair expires; this is a time-to-live (TTL) from the original publication date
	expire = 86400
	// refresh is the time after which an otherwise unaccessed bucket must be refreshe
	refresh = 3600
	// replicate is the interval between Kademlia replication events, when a node is required to publish its entire database
	replicate = 3600
	// republish is the time after which the original publisher must republish a key/value pair
	republish = 86400
)
