permit (
    principal,
    action in
        [Action::"GetOrder",
         Action::"UpdateOrder",
         Action::"CreateOrder",
         Action::"DeleteOrder"],
    resource
) when {
  principal.id == resource.user_id
};

