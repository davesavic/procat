-- name: RetrieveProductHierarchy :many
WITH RECURSIVE product_hierarchy AS (
    -- Base case: Start with the initial product
    SELECT
        p.id AS product_id,
        p.name AS product_name,
        p.description AS product_description,
        svf.name AS field_name,
        svf.data_type,
        pf.value,
        NULL::uuid AS parent_product_id,
        NULL::uuid AS composition_field_id,
        NULL::text AS composition_field_name,
        0 AS level
    FROM public.products p
    JOIN public.product_fields pf ON p.id = pf.product_id AND pf.deleted_at IS NULL
    JOIN public.schema_version_fields svf ON pf.schema_version_field_id = svf.id AND svf.deleted_at IS NULL
    WHERE p.id = @product_id
      AND p.deleted_at IS NULL

    UNION ALL

    SELECT
        cp.id AS product_id,
        cp.name AS product_name,
        cp.description AS product_description,
        svf_child.name AS field_name,
        svf_child.data_type,
        pf_child.value,
        pc.parent_product_id AS parent_product_id,
        pc.schema_version_field_id AS composition_field_id,
        svf_comp.name AS composition_field_name,
        ph.level + 1 AS level
    FROM product_hierarchy ph
    JOIN public.product_compositions pc ON pc.parent_product_id = ph.product_id AND pc.deleted_at IS NULL
    JOIN public.products cp ON cp.id = pc.child_product_id AND cp.deleted_at IS NULL
    JOIN public.product_fields pf_child ON pf_child.product_id = cp.id AND pf_child.deleted_at IS NULL
    JOIN public.schema_version_fields svf_child ON pf_child.schema_version_field_id = svf_child.id AND svf_child.deleted_at IS NULL
    JOIN public.schema_version_fields svf_comp ON pc.schema_version_field_id = svf_comp.id AND svf_comp.deleted_at IS NULL
)
SELECT
    product_id,
    product_name,
    product_description,
    field_name,
    data_type,
    value,
    parent_product_id,
    composition_field_id,
    composition_field_name,
    level
FROM product_hierarchy
ORDER BY level, product_id, field_name;
