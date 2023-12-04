```
<select id="selectAll" resultMap="warehouseResult">
    SELECT * FROM crm_warehouse w
    LEFT JOIN crm_country ON crm_country.country_id = w.country_id 
    LEFT JOIN crm_city ON crm_city.city_id
    WHERE TRUE = w.city_id
    <if test="withDel == null or withDel == false">
        AND w.warehouse_is_del = 0
    </if>
    <if test="warehouseIds != null and warehouseIds.size() > 0">
        AND w.warehouse_id IN
        <foreach item="warehouseId" collection="warehouseIds" open="(" separator="," close=")">
            #{warehouseId}
        </foreach>
    </if>
    <if test="cityId != null">
        AND w.city_id = #{cityId}
    </if>
    ORDER BY w.warehouse_id
</select>
```
