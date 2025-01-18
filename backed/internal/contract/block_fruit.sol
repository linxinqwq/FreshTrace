// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Marketplace {
    enum ProductStatus {
        Pending, // 未需要审核
        Await, // 等待中
        Approved, //已批准
        Rejected, // 已经拒绝
        ApprovedShelves //商品未上架
    }

    struct ProductionCondition {
        string date; // 日期，例如 "2024.3.1"
        string operation; // 操作，例如 "施肥"
        string imageId; // 相关操作的图片ID
    }

    struct Product {
        uint256 id;
        string imageId;
        string seller;
        uint256 stock;
        uint256 price;
        string uintGoods;
        string name;
        string description;
        uint256 sales;
        string origin;
        string kind;
        ProductStatus status;
    }

    struct ProductResult {
        uint256 id;
        string imageId;
        string seller;
        uint256 stock;
        uint256 price;
        string uintGoods;
        string name;
        string description;
        uint256 sales;
        string origin;
        string kind;
        ProductStatus status;
        ProductionCondition[] process;
    }

    // 使用产品ID映射到购买条件的映射
    mapping(uint256 => ProductionCondition[]) public process;

    mapping(string => PurchaseOrder[]) public userPurchases;



    struct PurchaseOrder {
        uint256 id;
        uint256 productId;
        uint256 quantity;
        uint256 blockNumber;
        string  createdTime;
        string  userAddress;
        bool status;
        string butName;
    }

    PurchaseOrder[] public totalOrders;
    Product[] public products;
    uint256 public nextProductId = 0;
    uint256 public nextOrderId =0;
    mapping(string => uint256[]) public purchases;
    mapping(string => uint256[]) public listings;

    // 事件定义
    event ProductAdded(uint256 productId, string seller);
    event ProductRemoved(uint256 productId, string seller);
    event ProductPurchased(uint256 productId, uint256 quantity, string buyer);
    event ProductReviewed(uint256 productId, ProductStatus status);

    // 添加产品
    function addProduct(
        string memory imageId,
        uint256 stock,
        uint256 price,
        string memory name,
        string memory uintGoods,
        string memory seller,
        string memory description,
        string memory origin,
        string memory kind
    ) public {
        products.push(
            Product(
                nextProductId,
                imageId,
                seller,
                stock,
                price,
                uintGoods,
                name,
                description,
                0,
                origin,
                kind,
                ProductStatus.Pending
            )
        );
        listings[seller].push(nextProductId);
        emit ProductAdded(nextProductId, seller);
        nextProductId++;
    }

    // 审核产品
    function reviewProduct(uint256 productId, ProductStatus status) public {
        products[productId].status = status;
        emit ProductReviewed(productId, status);
    }

    // 添加生产过程条件的函数
    function addProductionCondition(
        string memory username,
        uint256 productId,
        string memory date,
        string memory operation,
        string memory imageId,
        ProductStatus status
    ) public {
        // 获取商品信息
        Product storage product = products[productId];
        require(
            keccak256(bytes(username)) == keccak256(bytes(product.seller)),
            "dont seller"
        );

        require(product.status != ProductStatus.Approved, "sure status");
        // 更新状态
        product.status = status;

        // 创建一个新的ProductionCondition结构体实例
        ProductionCondition memory newCondition = ProductionCondition({
            date: date,
            operation: operation,
            imageId: imageId
        });

        // 将新的条件添加到商品的生产过程中
        process[productId].push(newCondition);

        // 可选：发出一个事件来记录生产条件的添加
        emit ProductionConditionAdded(productId, date, operation, imageId);
    }

    // 定义一个事件来记录生产条件的添加
    event ProductionConditionAdded(
        uint256 productId,
        string date,
        string operation,
        string imageId
    );

    // 修改购买产品函数，增加审核状态检查
    function buyProduct(
        uint256 productId,
        uint256 quantity,
        string memory buyer,
        string memory uAddress,
        string memory time
    ) public payable {
        Product storage product = products[productId];
        require(
            product.status == ProductStatus.Approved,
            "Product must be approved before purchase."
        );
        require(quantity <= product.stock, "Not enough stock.");

        product.stock -= quantity;
        product.sales += quantity;

        PurchaseOrder memory order = PurchaseOrder({
            id: nextOrderId,
            productId: productId,
            quantity: quantity,
            blockNumber: block.number,
            createdTime: time,
            userAddress: uAddress,
            status: false,
            butName: buyer
        });

        totalOrders.push(order);

        userPurchases[buyer].push(order);

        nextOrderId++;

        emit ProductPurchased(productId, quantity, buyer);
    }

    // 修改获取销售量前三的商品函数
    function getTopSellingProducts()
    public
    view
    returns (ProductResult[] memory)
    {
        ProductResult[] memory topSellingResults = new ProductResult[](3);
        uint256[] memory topSales = new uint256[](3);

        for (uint256 i = 0; i < products.length; i++) {
            if (products[i].status != ProductStatus.Approved) {
                continue; // 忽略未审核或审核不通过的商品
            }
            uint256 currentSales = products[i].sales;
            for (uint256 j = 0; j < topSellingResults.length; j++) {
                if (currentSales > topSales[j]) {
                    for (uint256 k = topSellingResults.length - 1; k > j; k--) {
                        topSellingResults[k] = topSellingResults[k - 1];
                        topSales[k] = topSales[k - 1];
                    }
                    // 创建ProductResult实例，包含产品信息和对应的生产过程
                    Product storage currentProduct = products[i];
                    ProductionCondition[] storage conditions = process[
                                    currentProduct.id
                        ];
                    topSellingResults[j] = ProductResult({
                        id: currentProduct.id,
                        imageId: currentProduct.imageId,
                        seller: currentProduct.seller,
                        stock: currentProduct.stock,
                        price: currentProduct.price,
                        uintGoods: currentProduct.uintGoods,
                        name: currentProduct.name,
                        description: currentProduct.description,
                        sales: currentSales,
                        origin: currentProduct.origin,
                        status: currentProduct.status,
                        kind: currentProduct.kind,
                        process: conditions
                    });
                    topSales[j] = currentSales;
                    break;
                }
            }
        }
        return topSellingResults;
    }

    // 其他函数保持不变

    // 下架产品
    function removeProduct(uint256 productId, string memory seller) public {
        require(
            keccak256(bytes(products[productId].seller)) ==
            keccak256(bytes(seller)),
            "Only the seller can remove this product."
        );
        delete products[productId];
        emit ProductRemoved(productId, seller);
    }

    // 修改产品信息
    function updateProduct(
        uint256 productId,
        uint256 newStock,
        uint256 newPrice,
        string memory newDescription,
        string memory name,
        string memory image,
        string memory seller,
        string memory origin,
        string memory kind,
        string memory uints
    ) public {
        Product storage product = products[productId];
        require(
            keccak256(bytes(product.seller)) == keccak256(bytes(seller)),
            "Only the seller can update this product."
        );

        product.stock = newStock;
        product.price = newPrice;
        product.description = newDescription;
        product.name = name;
        product.imageId = image;
        product.origin = origin;
        product.uintGoods = uints;
        product.kind = kind;
        // 无需修改审核状态，除非有专门逻辑要求

        // 可以考虑添加一个事件来记录产品信息更新
        // emit ProductUpdated(productId, product);
    }

    // 事件定义，用于记录产品信息被更新（可选）
    event ProductUpdated(uint256 productId, Product product);

    // 获取用户的购买记录（详细信息）
    function getUserPurchases(string memory user)
    public
    view
    returns (PurchaseOrder[] memory)
    {
        PurchaseOrder[] memory orders = userPurchases[user];
        return orders;
    }

    // 获取卖家发布的商品列表（详细信息）
    function getSellerListings(string memory seller)
    public
    view
    returns (ProductResult[] memory)
    {
        uint256[] memory listingIds = listings[seller];
        ProductResult[] memory sellerProductResults = new ProductResult[](
            listingIds.length
        );

        for (uint256 i = 0; i < listingIds.length; i++) {
            Product storage currentProduct = products[listingIds[i]];
            ProductionCondition[] storage conditions = process[
                            currentProduct.id
                ];

            sellerProductResults[i] = ProductResult({
                id: currentProduct.id,
                imageId: currentProduct.imageId,
                seller: currentProduct.seller,
                stock: currentProduct.stock,
                price: currentProduct.price,
                uintGoods: currentProduct.uintGoods,
                name: currentProduct.name,
                description: currentProduct.description,
                sales: currentProduct.sales,
                origin: currentProduct.origin,
                kind: currentProduct.kind,
                status: currentProduct.status,
                process: conditions
            });
        }

        return sellerProductResults;
    }

    // 获取待审核的商品列表，包括生产过程信息
    function getPendingProducts() public view returns (ProductResult[] memory) {
        uint256 pendingCount = 0;
        // 首先，计算待审核商品的数量
        for (uint256 i = 0; i < products.length; i++) {
            if (products[i].status == ProductStatus.Await) {
                pendingCount++;
            }
        }

        // 根据待审核商品数量创建一个新数组
        ProductResult[] memory pendingProductResults = new ProductResult[](
            pendingCount
        );
        uint256 currentIndex = 0;

        // 再次遍历商品，为待审核的商品创建ProductResult实例，并添加到数组中
        for (uint256 i = 0; i < products.length; i++) {
            if (products[i].status == ProductStatus.Await) {
                Product storage currentProduct = products[i];
                ProductionCondition[] storage conditions = process[
                                currentProduct.id
                    ];

                pendingProductResults[currentIndex] = ProductResult({
                    id: currentProduct.id,
                    imageId: currentProduct.imageId,
                    seller: currentProduct.seller,
                    stock: currentProduct.stock,
                    price: currentProduct.price,
                    uintGoods: currentProduct.uintGoods,
                    name: currentProduct.name,
                    description: currentProduct.description,
                    sales: currentProduct.sales,
                    origin: currentProduct.origin,
                    kind: currentProduct.kind,
                    status: currentProduct.status,
                    process: conditions
                });
                currentIndex++;
            }
        }

        return pendingProductResults;
    }

    // 获取商品详情
    // 获取商品详情
    function getProductDetail(uint256 id)
    public
    view
    returns (ProductResult memory)
    {
        // 获取商品信息
        Product storage product = products[id];

        // 获取商品的生产过程信息
        ProductionCondition[] storage conditions = process[id];

        // 构造 ProductResult 结构体
        ProductResult memory result = ProductResult({
            id: product.id,
            imageId: product.imageId,
            seller: product.seller,
            stock: product.stock,
            price: product.price,
            uintGoods: product.uintGoods,
            name: product.name,
            description: product.description,
            sales: product.sales,
            origin: product.origin,
            kind: product.kind,
            status: product.status,
            process: conditions
        });

        return result;
    }

    // 根据商品种类查询商品
    function getProductsByKind(string memory _kind)
    public
    view
    returns (ProductResult[] memory)
    {
        uint256 count = 0;
        // 首先，计算符合条件的商品数量
        for (uint256 i = 0; i < products.length; i++) {
            if (keccak256(bytes(products[i].kind)) == keccak256(bytes(_kind))) {
                count++;
            }
        }

        // 根据符合条件的商品数量创建一个新数组
        ProductResult[] memory result = new ProductResult[](count);
        uint256 currentIndex = 0;

        // 再次遍历商品数组，将符合条件的商品添加到新数组中
        for (uint256 i = 0; i < products.length; i++) {
            if (keccak256(bytes(products[i].kind)) == keccak256(bytes(_kind))) {
                // 获取商品的生产过程信息
                ProductionCondition[] storage conditions = process[
                                    products[i].id
                    ];

                // 构造 ProductResult 结构体
                result[currentIndex] = ProductResult({
                    id: products[i].id,
                    imageId: products[i].imageId,
                    seller: products[i].seller,
                    stock: products[i].stock,
                    price: products[i].price,
                    uintGoods: products[i].uintGoods,
                    name: products[i].name,
                    description: products[i].description,
                    sales: products[i].sales,
                    origin: products[i].origin,
                    kind: products[i].kind,
                    status: products[i].status,
                    process: conditions
                });

                currentIndex++;
            }
        }

        return result;
    }

    // 方法来设置订单状态，并同步更新userPurchases中的状态
    function setOrderStatus(uint256 orderIndex, bool newStatus) public {
        require(orderIndex < totalOrders.length, "Invalid order index.");

        // 更新totalOrders中的订单状态
        PurchaseOrder storage order = totalOrders[orderIndex];
        order.status = newStatus;

        // 同步更新userPurchases中的订单状态
        updatePurchaseOrderInUserPurchases(order.butName, order.id, newStatus);

        emit OrderStatusUpdated(order.id, newStatus);
    }

    // 定义一个事件来记录订单状态的更新
    event OrderStatusUpdated(uint256 orderIndex, bool newStatus);

    // 辅助函数，用于在userPurchases中更新订单状态
    function updatePurchaseOrderInUserPurchases(string memory buyerName, uint256 orderId, bool newStatus) private {
        PurchaseOrder[] storage orders = userPurchases[buyerName];
        for (uint256 i = 0; i < orders.length; i++) {
            if (orders[i].id == orderId) {
                orders[i].status = newStatus;
                break;
            }
        }
    }
}
