package inventory

//type Client struct {
//	conn   *grpc.ClientConn
//	client inventorypb.InventoryServiceClient
//}
//
//func NewClient(cfg *config.Config) (*Client, error) {
//	conn, err := grpc.Dial(cfg.InventoryServiceGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		return nil, fmt.Errorf("failed to connect to inventory service: %v", err)
//	}
//
//	c := inventorypb.NewInventoryServiceClient(conn)
//
//	return &Client{
//		conn:   conn,
//		client: c,
//	}, nil
//}
//
//func (c *Client) Close() error {
//	return c.conn.Close()
//}
//
//func (c *Client) GetProduct(ctx context.Context, productID uint) (*inventorypb.ProductResponse, error) {
//	resp, err := c.client.GetProduct(ctx, &inventorypb.GetProductRequest{
//		ProductId: uint32(productID),
//	})
//	if err != nil {
//		return nil, fmt.Errorf("failed to get product: %v", err)
//	}
//	return resp, nil
//}
//
//func (c *Client) UpdateProductStock(ctx context.Context, productID uint, quantity int) error {
//	_, err := c.client.UpdateProductStock(ctx, &inventorypb.UpdateProductStockRequest{
//		ProductId: uint32(productID),
//		Stock:     int32(quantity),
//	})
//	if err != nil {
//		return fmt.Errorf("failed to update stock: %v", err)
//	}
//	return nil
//}
